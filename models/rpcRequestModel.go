package models

import (
	"gorm.io/gorm"
)

type RpcRequest struct {
	gorm.Model
	AccountId      uint
	MethodName     string
	RequestBody    string
	ResponseStatus uint
	ResponseBody   string
	Chain          string
	Network        string
	IpAddress      string
	EndpointID     uint
	IsTest         bool `gorm:"default:false"`
}

func (r *RpcRequest) Successful() bool {
	return r.ResponseStatus >= 200 && r.ResponseStatus <= 299
}

func (r *RpcRequest) PrettyCreatedAt() string {
	return r.CreatedAt.Format("2006-01-02 15:04:05") // or any other format
}
