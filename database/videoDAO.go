package database

import (
	"errors"

	"github.com/asciiflix/server/model"
)

//Create video
func CreateVideo(video model.Video) error {

	result := global_db.Create(&video)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

//Get video by id
func GetVideo(videoId string) (*model.Video, error) {
	var video model.Video
	result := global_db.Preload("Comments").Where("uuid = ?", videoId).First(&video)

	if result.Error != nil {
		return nil, result.Error
	}
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
	result := global_db.Where("uuid = ?", updateVideo.UUID).First(&videoToUpdate)
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

//Delete video by uuid
func DeleteVideo(videoId string, userId string) error {
	//Check if video exists and belongs to user
	result := global_db.Where("uuid = ? AND user_id = ?", videoId, userId).First(&model.Video{})
	if result.Error != nil {
		return result.Error

	}
	global_db.Where("uuid = ? AND user_id = ?", videoId, userId).Delete(&model.Video{})
	return nil
}

func GetContentID(videoUUID string) (string, error) {
	var video model.Video
	result := global_db.Where("uuid = ?", videoUUID).First(&video)

	if result.Error != nil {
		return "", result.Error
	}
	return video.VideoContentID, nil
}
