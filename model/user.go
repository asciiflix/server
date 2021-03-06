package model

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
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
	Verified    bool
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
	Verified    bool
	Videos      []VideoPublic
	Comments    []Comment
	Likes       []Like
}

//UserDetails for Public Endpoints (Profile Page etc)
type UserDetailsPublic struct {
	UserID      uint
	Name        string
	Description string
	Picture_ID  string
	Videos      []VideoPublic
}

//Login Struct
type UserLogin struct {
	Email    string
	Password string
}

//VerificationCode struct
type VerificationCode struct {
	ID     uint `gorm:"primaryKey"`
	UserID uint
	Code   string
	Expiry time.Time
}

//Login Struct
type JwtBlacklist struct {
	ID  uint `gorm:"primaryKey"`
	Jwt string
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
	for _, videos := range user.Videos {
		publicUser.Videos = append(publicUser.Videos, GetPublicVideo(videos))
	}
	return publicUser
}

func (user User) GetPrivateUser() UserDetailsPrivate {
	var privateUser UserDetailsPrivate
	privateUser.Name = user.Name
	privateUser.UserID = user.ID
	privateUser.Email = user.Email
	privateUser.Description = user.Description
	privateUser.Picture_ID = user.Picture_ID
	privateUser.Comments = user.Comments
	privateUser.Likes = user.Likes
	privateUser.Verified = user.Verified
	for _, videos := range user.Videos {
		privateUser.Videos = append(privateUser.Videos, GetPublicVideo(videos))
	}
	return privateUser
}
