package commandEvaluators

import (
	"github.com/byuoitav/av-api-rpc/base"
	avbase "github.com/byuoitav/av-api/base"
)

//Volumeupdefault the the default volume up command evaluator for rpc volume up commands.
type MuteDefault struct{}

func (v *MuteDefault) Evaluate(req base.RPCRequest) ([]avbase.ActionStructure, error) {
	var CommandName = "Mute"
	var EvaluatorName = "MuteDefault-RPC"
	return generatePassthroughCommand(req, CommandName, EvaluatorName)
}

func (v *MuteDefault) Validate(toEval avbase.ActionStructure) error {
	var CommandName = "Mute"
	return checkForCommandInDevice(toEval, CommandName)
}

func (v *MuteDefault) GetIncompatibleCommands() []string {
	return []string{}
}
