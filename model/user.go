package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Email    string
	Password string
}

type Login struct {
	Email    string
	Password string
}
