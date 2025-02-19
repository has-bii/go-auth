package models

import "github.com/google/uuid"

type User struct {
	Base
	Name     string `json:"name"`
	Email    string `gorm:"unique;not null" json:"email"`
	Password string
}

type UserDTO struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Email string    `json:"email"`
}

func (UserDTO) TableName() string {
	return "users"
}
