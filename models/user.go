package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type User struct {
	ID         uuid.UUID `json:"id" gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Name       string    `json:"name" `
	Address    string    `json:"andress" `
	CreateddAt time.Time `json:"created" `
	UpdatedAt  time.Time `json:"updated" `
}
