package types

import "time"

type RpcMethodCallsByDay struct {
	MethodName    string    `json:methodName`
	TruncatedDate time.Time `json:day`
	Count         int       `json:count`
}
