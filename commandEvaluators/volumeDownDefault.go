package commandEvaluators

import (
	"github.com/byuoitav/av-api-rpc/base"
	avbase "github.com/byuoitav/av-api/base"
)

type VolumeDownDefault struct{}

func (v *VolumeDownDefault) Evaluate(req base.RPCRequest) ([]avbase.ActionStructure, error) {

	var EvaluatorName = "VolumeDownDefault-RPC"
	var CommandName = "VolumeDown"
	return generatePassthroughCommand(req, CommandName, EvaluatorName)
}

func (v *VolumeDownDefault) Validate(toEval avbase.ActionStructure) error {

	var CommandName = "VolumeDown"
	return checkForCommandInDevice(toEval, CommandName)
}

func (v *VolumeDownDefault) GetIncompatibleCommands() []string {
	return []string{}
}
