package models

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	QuicknodeID     string `gorm:"index"`
	PlanSlug        string
	DeprovisionedAt datatypes.Time
	IsTest          bool `gorm:"default:false"`
}
