package model

import (
	"time"

	"github.com/gofrs/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gorm.io/gorm"
)

type Video struct {
	gorm.Model
	UUID           uuid.UUID //`gorm:"primary_key; unique; type:uuid; column:id;"`
	VideoContentID string
	Title          string
	Description    string
	UploadDate     time.Time
	Views          int
	UserID         uint
	Comments       []Comment `gorm:"ForeignKey:VideoID"`
	Likes          []Like    `gorm:"ForeignKey:VideoID"`
}

type VideoContent struct {
	ObjectID primitive.ObjectID     `bson:"_id,omitempty"`
	Video    map[string]interface{} `bson:"video"`
}

type VideoFull struct {
	VideoStats   Video
	VideoContent VideoContent
}
