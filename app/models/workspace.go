package models

import "github.com/google/uuid"

type Workspace struct {
	Base
	Name    string    `json:"name" gorm:"not null"`
	OwnerID uuid.UUID `gorm:"not null;primaryKey;type:uuid" json:"owner_id"`

	Owner   User   `gorm:"foreignKey:ID;references:OwnerID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"owner,omitempty"`
	Members []User `gorm:"many2many:workspace_users;joinForeignKey:WorkspaceID;joinReferences:UserID" json:"members,omitempty"`
}

type WorkspaceUser struct {
	WorkspaceID uuid.UUID `gorm:"type:uuid;primaryKey;" json:"workspace_id"`
	UserID      uuid.UUID `gorm:"type:uuid;primaryKey;" json:"user_id"`
}

type WorkspaceDTO struct {
	ID      uuid.UUID `json:"id"`
	Name    string    `json:"name"`
	OwnerID uuid.UUID `json:"owner_id"`
	Owner   UserDTO   `json:"owner"`
}
