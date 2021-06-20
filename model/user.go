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
	Videos      []Video   `gorm:"ForeignKey:UserID"`
	Comments    []Comment `gorm:"ForeignKey:UserID"`
	Likes       []Like    `gorm:"ForeignKey:UserID"`
}

//UserDetails for Private Endpoints (Settings etc.)
type UserDetailsPrivate struct {
	UserID      uint
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
	UserID      uint
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

func (user User) GetPublicUser() UserDetailsPublic {
	var publicUser UserDetailsPublic
	publicUser.Name = user.Name
	publicUser.UserID = user.ID
	publicUser.Description = user.Description
	publicUser.Picture_ID = user.Picture_ID
	publicUser.Videos = user.Videos

	return publicUser
}

func (user User) GetPrivateUser() UserDetailsPrivate {
	var privateUser UserDetailsPrivate
	privateUser.Name = user.Name
	privateUser.UserID = user.ID
	privateUser.Email = user.Email
	privateUser.Description = user.Description
	privateUser.Picture_ID = user.Picture_ID
	privateUser.Videos = user.Videos
	privateUser.Comments = user.Comments
	privateUser.Likes = user.Likes

	return privateUser
}
