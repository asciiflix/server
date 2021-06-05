package model

import "gorm.io/gorm"

type Video struct {
	gorm.Model
	Name        string
	Description string
	UserID      uint
	Comments    []Comment
	Likes       []Like
}
