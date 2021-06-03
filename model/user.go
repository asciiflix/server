package model

import (
	"github.com/dgrijalva/jwt-go"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Email    string
	Password string
}

//Login Struct
type UserLogin struct {
	Email    string
	Password string
}

//Struct for Users JWT Token
type UserClaim struct {
	jwt.StandardClaims
	User_ID    uint
	User_email string
}
