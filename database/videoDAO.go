package database

import (
	"errors"
	"fmt"

	"github.com/asciiflix/server/model"
	"github.com/gofrs/uuid"
)

//Create video
func CreateVideo(video model.Video) error {
	video.UUID, _ = uuid.NewV4()
	result := global_db.Create(&video)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

//Get video by id
func GetVideo(videoId string) (*model.Video, error) {
	var video model.Video
	result := global_db.Preload("Comments").Where("id = ?", videoId).First(&video)

	if result.Error != nil {
		return nil, result.Error
	}
	fmt.Println(video)
	return &video, nil
}

//Get all Videos
func GetVideos() (*[]model.Video, error) {
	var videos []model.Video
	result := global_db.Find(&videos)
	if result.Error != nil {
		return nil, result.Error
	}
	return &videos, nil
}

//Update video by id
func UpdateVideo(updateVideo model.Video) error {
	//Check if Video exists by ID
	var videoToUpdate model.Video
	result := global_db.Where("id = ?", updateVideo.ID).First(&videoToUpdate)
	if result.Error != nil {
		return errors.New("video does not exist")
	}

	//Updates Values in Database
	result = global_db.Model(&videoToUpdate).Updates(updateVideo)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

//Delete video by id
func DeleteVideo(videoId string, userId string) error {
	result := global_db.Where("id = ?", videoId).Where("user_id = ?", userId).Delete(&model.Video{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}
