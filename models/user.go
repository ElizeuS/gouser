package models

import (
	"time"
)

type User struct {
	ID           string    `json:"id" gorm:"default:uuid_generate_v4()"`
	Name         string    `json:"name" `
	Address      string    `json:"andress" `
	RegisteredAt time.Time `json:"created" `
	UpdatedAt    time.Time `json:"updated" `
}
