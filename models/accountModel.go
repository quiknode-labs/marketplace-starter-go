package models

import (
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	QuicknodeID string
	Plan        string
	IsTest      bool `gorm:"default:false"`
}
