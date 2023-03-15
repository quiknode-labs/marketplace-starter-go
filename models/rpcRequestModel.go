package models

import (
	"gorm.io/gorm"
)

type RpcRequest struct {
	gorm.Model
	QuicknodeID string
	Chain       string
	Network     string
	Method      string
	RequestID   string
	Version     string
	IsTest      bool `gorm:"default:false"`
}
