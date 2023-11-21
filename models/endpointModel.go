package models

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Endpoint struct {
	gorm.Model
	AccountID       uint `gorm:"index"`
	QuicknodeID     string
	WssUrl          string
	HttpUrl         string
	Chain           string
	Network         string
	DeprovisionedAt datatypes.Time
	IsTest          bool `gorm:"default:false"`
}
