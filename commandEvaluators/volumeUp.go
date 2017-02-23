package main

import (
	"github.com/byuoitav/av-api-rpc/base"
	avbase "github.com/byuoitav/av-api/base"
)

//Volumeupdefault the the default volume up command evaluator for rpc volume up commands.
type VolumeUpDefault struct{}

var CommandName = "VolumeUp"
var EvaluatorName = "VolumeUpDefault-RPC"

func (v *VolumeUpDefault) Evaluate(req base.RPCRequest) ([]avbase.ActionStructure, error) {
	return generatePassthroughCommand(req, CommandName, EvaluatorName)
}

func (v *VolumeUpDefault) Validate(toEval avbase.ActionStructure) error {
	return checkForCommandInDevice(toEval, CommandName)
}

func (v *VolumeUpDefault) GetIncompatibleCommands() []string {
	return []string{}
}
