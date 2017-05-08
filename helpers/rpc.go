package helpers

import (
	"fmt"
	"log"
	"regexp"

	apibase "github.com/byuoitav/av-api/base"

	"github.com/byuoitav/av-api-rpc/base"
	"github.com/byuoitav/av-api-rpc/commandEvaluators"
	"github.com/byuoitav/av-api-rpc/reconcilers"
	apiEvals "github.com/byuoitav/av-api/commandevaluators"
	"github.com/byuoitav/av-api/dbo"
)

func RunCommands(request base.RPCRequest) ([]apiEvals.CommandExecutionReporting, error) {
	log.Printf("Getting room %s from the Database", request.Room)
	//get the room from the databse
	room, err := dbo.GetRoomByInfo(request.Building, request.Room)
	if err != nil {
		log.Printf("Error getting data. ERROR: %s", err.Error())
		return []apiEvals.CommandExecutionReporting{}, nil
	}

	evaulatorMap := commandEvaluators.Init()

	ActionList := []apibase.ActionStructure{}

	re := regexp.MustCompile(".*-RPC$")

	log.Printf("Beginning evaluation run...")
	for _, evaluator := range room.Configuration.Evaluators {
		fmt.Printf("Command Key: %s\n", evaluator.EvaluatorKey)
		if re.MatchString(evaluator.EvaluatorKey) { //Check to make sure that it's an RPC evaulator
			log.Printf("Found evaluator: %s, evaluating", evaluator.EvaluatorKey)
			if evalStruct, ok := evaulatorMap[evaluator.EvaluatorKey]; ok {
				tempList, err := evalStruct.Evaluate(request)
				if err != nil {
					log.Printf("Error generating actions for %s. ERROR: %s", evaluator.EvaluatorKey, err.Error())
					return []apiEvals.CommandExecutionReporting{}, err
				}
				log.Printf("%v events generated", len(tempList))
				log.Printf("Starting validation run...")
				for _, action := range tempList {
					err := evalStruct.Validate(action)

					if err != nil {
						log.Printf("Error generating commands. Command %s is an invalid command for %s. Error: %s",
							action.Action,
							action.Device.GetFullName(),
							err.Error())

						return []apiEvals.CommandExecutionReporting{}, err
					}
				}
				ActionList = append(ActionList, tempList...)
				log.Printf("Evaulator finished.")
			}
		}
	}
	log.Printf("Evaluation run finished")
	log.Printf("Beginning reconciliation...")
	reconcilerMap := reconcilers.Init()
	//get the room configuration from the list.
	if reconciler, ok := reconcilerMap[room.Configuration.RoomKey+"-RPC"]; ok {
		ActionList, err = reconciler.Reconcile(ActionList)
		if err != nil {
			log.Printf("Error reconciling the commands. Error.: %s", err.Error())
			return []apiEvals.CommandExecutionReporting{}, err
		}
	}
	log.Printf("Reconciliation finished.")

	return apiEvals.ExecuteActions(ActionList)
}
