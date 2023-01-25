package rpc

import(
  "encoding/json"
)

type RpcRequest struct {
  ID      interface{}       `json:"id"`
  JsonRpc string            `json:"jsonrpc"`
  Method  string            `json:"method"`
  Params  []json.RawMessage `json:"params"`
}
