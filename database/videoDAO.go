package database

import (
	"github.com/asciiflix/server/model"
	"github.com/gofrs/uuid"
)

//Testing for UUID-GEN
func CreateVideo() {
	var video model.Video

	video.ID, _ = uuid.NewV4()
	video.Title = "Test Title"
	video.Description = "Test Desc"

	global_db.Create(&video)
}
