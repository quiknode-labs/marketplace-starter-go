package types

type JsonRpcErrorResponse struct {
	ID      string `json:"id"`
	Error   Error  `json:"error"`
	Jsonrpc string `json:"jsonrpc"`
}

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
