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

type VideoPublic struct {
	UUID        uuid.UUID
	Title       string
	Description string
	UploadDate  time.Time
	Views       int
	Likes       int
	UserID      uint
	//Comments    []Comment //Comments need own endpoint
}

type VideoContent struct {
	ObjectID primitive.ObjectID     `bson:"_id,omitempty"`
	Video    map[string]interface{} `bson:"video"`
}

type VideoFull struct {
	VideoStats   Video
	VideoContent VideoContent
}

func GetPublicVideo(video Video) VideoPublic {
	return VideoPublic{
		UUID:        video.UUID,
		Title:       video.Title,
		Description: video.Description,
		UploadDate:  video.UploadDate,
		Views:       video.Views,
		Likes:       len(video.Likes),
		UserID:      video.UserID,
		//Comments:    video.Comments,
	}
}
