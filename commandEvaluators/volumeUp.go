package commandEvaluators

import (
	"github.com/byuoitav/av-api-rpc/base"
	avbase "github.com/byuoitav/av-api/base"
)

//Volumeupdefault the the default volume up command evaluator for rpc volume up commands.
type VolumeUpDefault struct{}

func (v *VolumeUpDefault) Evaluate(req base.RPCRequest) ([]avbase.ActionStructure, error) {
	var CommandName = "VolumeUp"
	var EvaluatorName = "VolumeUpDefault-RPC"
	return generatePassthroughCommand(req, CommandName, EvaluatorName)
}

func (v *VolumeUpDefault) Validate(toEval avbase.ActionStructure) error {
	var CommandName = "VolumeUp"
	return checkForCommandInDevice(toEval, CommandName)
}

func (v *VolumeUpDefault) GetIncompatibleCommands() []string {
	return []string{}
}
