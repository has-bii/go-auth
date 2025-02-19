package models

import "github.com/google/uuid"

type Card struct {
	Base
	Name        string    `json:"name" gorm:"not null"`
	Description string    `json:"description"`
	Order       int       `json:"order" gorm:"not null"`
	ListID      uuid.UUID `gorm:"type:uuid;primaryKey" json:"list_id"`
}
