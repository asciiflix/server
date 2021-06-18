package model

import (
	"github.com/dgrijalva/jwt-go"
	"gorm.io/gorm"
)

//User struct for DB and backend stuff
type User struct {
	gorm.Model
	Name        string
	Email       string
	Password    string
	Picture_ID  string
	Description string
	Videos      []Video
	Comments    []Comment
	Likes       []Like
}

//UserDetails fpr Private Endpoints (Settings etc.)
type UserDetailsPrivate struct {
	Name        string
	Email       string
	Description string
	Picture_ID  string
	Videos      []Video
	Comments    []Comment
	Likes       []Like
}

//UserDetails for Public Endpoints (Profile Page etc)
type UserDetailsPublic struct {
	Name        string
	Description string
	Picture_ID  string
	Videos      []Video
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

//Todo Parse function public
//Todo parse function private
