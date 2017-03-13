package commandEvaluators

import (
	"github.com/byuoitav/av-api-rpc/base"
	avbase "github.com/byuoitav/av-api/base"
)

//Volumeupdefault the the default volume up command evaluator for rpc volume up commands.
type ToggleMuteDefault struct{}

func (v *ToggleMuteDefault) Evaluate(req base.RPCRequest) ([]avbase.ActionStructure, error) {
	var CommandName = "ToggleMute"
	var EvaluatorName = "ToggleMute-RPC"
	return generatePassthroughCommand(req, CommandName, EvaluatorName)
}

func (v *ToggleMuteDefault) Validate(toEval avbase.ActionStructure) error {
	var CommandName = "ToggleMute"
	return checkForCommandInDevice(toEval, CommandName)
}

func (v *ToggleMuteDefault) GetIncompatibleCommands() []string {
	return []string{}
}
