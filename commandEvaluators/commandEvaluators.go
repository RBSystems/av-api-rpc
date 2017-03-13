package commandEvaluators

import (
	"errors"
	"fmt"
	"log"

	"github.com/byuoitav/av-api-rpc/base"
	avbase "github.com/byuoitav/av-api/base"
	avdbo "github.com/byuoitav/av-api/dbo"
)

//CommandExecutionReporting is a struct we use to keep track of command execution
//for reporting to the user.
type CommandExecutionReporting struct {
	Success bool   `json:"success"`
	Action  string `json:"action"`
	Device  string `json:"device"`
	Err     string `json:"error,omitempty"`
}

var commandMapInitialized = false

/*
CommandEvaluator is an interface that must be implemented for each command to be
evaluated.
*/
type CommandEvaluator interface {
	/*
		Evaluate takes the RPC request (from the PUT body) and generates an action struct.

		While this will usually be a simple device lookup and translation - there are cases
		where parameters will need to be fetched from the databse
	*/
	Evaluate(base.RPCRequest) ([]avbase.ActionStructure, error)
	/*
		  Validate takes an action structure (for the command) and validates
			that the device and parameter are valid for the command.
	*/
	Validate(avbase.ActionStructure) error
	/*
			   GetIncompatableActions returns a list of commands that are incompatable
		     with this one (i.e. 'standby' and 'power on', or 'mute' and 'volume up')
	*/
	GetIncompatibleCommands() []string
}

/*
Need to pull in the list of commands, look through them for the relevant commands, then generate a new command for them.

This may take the same form as the av-api.ActionStructure - in fact I'm almost certain that it will. We can reuse that package.
*/

/*
generatePassthroughCommand can be a generic evaluator for each command when that command requires no parameter mapping from the databse.
*/
func generatePassthroughCommand(req base.RPCRequest, CommandName string, generatingEvaluator string) ([]avbase.ActionStructure, error) {
	toreturn := []avbase.ActionStructure{}
	log.Printf("Starting evaluation for %s", CommandName)
	for _, device := range req.RPCDevices {
		log.Printf("Looking in device %s for applicable commands", device.Name)
		for _, command := range device.Commands {
			if command.Name == CommandName {

				log.Printf("Command found for device %s", device.Name)

				databaseDevice, err := avdbo.GetDeviceByName(req.Building, req.Room, device.Name)
				if err != nil {
					log.Printf("Error getting device from databse. ERROR: %v", err.Error())
					return toreturn, err
				}

				params := make(map[string]string)

				for _, param := range command.Parameters {
					params[param.Name] = param.Value
				}

				toreturn = append(toreturn, avbase.ActionStructure{
					Action:              CommandName,
					GeneratingEvaluator: generatingEvaluator,
					Device:              databaseDevice,
					Parameters:          params,
					DeviceSpecific:      true,
				})
				log.Printf("Event Generated")
			}
		}
	}
	return toreturn, nil

}

func checkForCommandInDevice(toEval avbase.ActionStructure, commandName string) error {
	log.Printf("Validating that command %s is a valid command for device %s", toEval.Action, toEval.Device.GetFullName())
	for _, cmd := range toEval.Device.Commands {
		if cmd.Name == commandName {
			log.Printf("Command validated.")
			return nil
		}
	}
	message := fmt.Sprintf("There is no command %s for the device %s", commandName, toEval.Device.GetFullName())
	log.Printf("%s", message)
	return errors.New(message)

}

var CommandMap = make(map[string]CommandEvaluator)

/*
Init simply initializes the map of RPC command evaluators
*/
func Init() map[string]CommandEvaluator {
	if !commandMapInitialized {

		CommandMap["VolumeUpDefault-RPC"] = &VolumeUpDefault{}
		CommandMap["VolumeDownDefault-RPC"] = &VolumeDownDefault{}
		CommandMap["ToggleMute-RPC"] = &MuteDefault{}
		commandMapInitialized = true
	}

	return CommandMap
}
