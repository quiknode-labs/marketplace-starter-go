package models

import(
  "gorm.io/gorm"
)

type Endpoint struct {
  gorm.Model
  AccountID      uint
  QuicknodeID    string
  WssUrl         string
  HttpUrl        string
  Chain          string
  Network        string
}
