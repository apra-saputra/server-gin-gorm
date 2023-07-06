package model

import (
	"restapi/modules/response"

	"gorm.io/gorm"
)

type RoleType string

const (
	Users RoleType = "user"
	Admin RoleType = "admin"
	Dev   RoleType = "dev"
)

type User struct {
	gorm.Model
	ID       uint `gorm:"primaryKey"`
	Username string
	Email    string `gorm:"unique"`
	Address  string
	Otp      string
	Role     RoleType `gorm:"default:'user'"`
	Tasks    []Task
}

func (u *User) GetUser() response.UserResponse {
	return response.UserResponse{
		Id:       u.ID,
		Email:    u.Email,
		Username: u.Username,
		Address:  u.Address,
	}
}
