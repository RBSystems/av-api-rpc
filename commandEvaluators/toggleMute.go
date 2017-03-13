package commandEvaluators

import (
	"github.com/byuoitav/av-api-rpc/base"
	avbase "github.com/byuoitav/av-api/base"
)

//Volumeupdefault the the default volume up command evaluator for rpc volume up commands.
type UnMuteDefault struct{}

func (v *UnMuteDefault) Evaluate(req base.RPCRequest) ([]avbase.ActionStructure, error) {
	var CommandName = "ToggleMute"
	var EvaluatorName = "ToggleMute-RPC"
	return generatePassthroughCommand(req, CommandName, EvaluatorName)
}

func (v *UnMuteDefault) Validate(toEval avbase.ActionStructure) error {
	var CommandName = "ToggleMute"
	return checkForCommandInDevice(toEval, CommandName)
}

func (v *UnMuteDefault) GetIncompatibleCommands() []string {
	return []string{}
}
