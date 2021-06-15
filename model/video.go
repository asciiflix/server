package model

import (
	"github.com/gofrs/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gorm.io/gorm"
)

type Video struct {
	gorm.Model
	Name        string
	Description string
	UserID      uint
	Comments    []Comment
	Likes       []Like
}

type VideoStats struct {
	ID          uuid.UUID `gorm:"primary_key; unique; type:uuid; column:id;"`
	ObjectID    string
	Title       string
	Description string
}

type VideoContent struct {
	ID    primitive.ObjectID     `bson:"_id,omitempty"`
	Video map[string]interface{} `bson:"video"`
}
