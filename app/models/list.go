package models

import "github.com/google/uuid"

type List struct {
	Base
	Name        string    `json:"name" gorm:"not null"`
	Order       int       `json:"order" gorm:"not null"`
	WorkspaceID uuid.UUID `gorm:"type:uuid;primaryKey;not null" json:"workspace_id"`

	Cards []Card `json:"cards"`
}
