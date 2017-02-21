package base

//RPCRequest and it's children structs represent the expected put body for this command.
type RPCRequest struct {
	RPCDevices []RPCDevice `json:"rpcDevices"`
}

//RPCDevice .
type RPCDevice struct {
	Name     string       `json:"name"`
	Commands []RPCCommand `json:"commands"`
}

//RPCCommand .
type RPCCommand struct {
	Name       string         `json:"name"`
	Parameters []RPCParameter `json:"parameters,omitempty"`
	Success    bool           `json:"success,omitempty"`
}

//RPCParameter .
type RPCParameter struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
