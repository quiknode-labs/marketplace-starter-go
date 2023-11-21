package types

type RpcRequestPayload struct {
	ID         string      `json:"id"` // ID can be an integer or a string
	MethodName string      `json:"method"`
	Params     interface{} `json:"params"`
}
