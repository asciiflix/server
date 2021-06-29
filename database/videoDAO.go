package database

import (
	"errors"
	"fmt"

	"github.com/asciiflix/server/model"
	"gorm.io/gorm/clause"
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
	result := global_db.Preload("Comments").Preload("Likes").Where("uuid = ?", videoId).First(&video)

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

func GetRecomendations(limit int) (*[]model.Video, error) {
	var videos []model.Video
	result := global_db.Limit(limit).Clauses(clause.OrderBy{Expression: clause.Expr{SQL: "RANDOM()"}}).Find(&videos)
	if result.Error != nil {
		return nil, result.Error
	}
	return &videos, nil
}

func GetRecomendationsForUser(limit int, uuid uint) (*[]model.Video, error) {
	var videos []model.Video
	result := global_db.Limit(limit).Where("user_id != ?", uuid).Clauses(clause.OrderBy{Expression: clause.Expr{SQL: "RANDOM()"}}).Find(&videos)
	if result.Error != nil {
		return nil, result.Error
	}
	return &videos, nil
}

//Get all Videos from one user
func GetVideosFromUser(userID string) (*[]model.Video, error) {
	var videos []model.Video
	result := global_db.Where("user_id = ?", userID).Find(&videos)
	if result.Error != nil {
		return nil, result.Error
	}
	return &videos, nil
}

//Update video by id
func UpdateVideo(updateVideo model.Video, userId uint) error {
	//Check if Video exists by ID
	var videoToUpdate model.Video
	result := global_db.Where("uuid = ?", updateVideo.UUID).First(&videoToUpdate)
	if result.Error != nil {
		return errors.New("video does not exist")
	}

	if userId != videoToUpdate.UserID {
		fmt.Println(videoToUpdate)
		return errors.New("user does not match")
	}

	//Updates Values in Database
	result = global_db.Model(&videoToUpdate).Updates(updateVideo)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

//Delete video by uuid
func DeleteVideo(videoId string, userId uint) error {
	//Check if video exists and belongs to user
	video, err := GetVideo(videoId)
	if err != nil {
		fmt.Println(videoId)
		return err
	}
	if userId != video.UserID {
		fmt.Println(videoId)
		return errors.New("user does not match")
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

func GetVideoID(videoUUID string) (uint, error) {
	var video model.Video
	result := global_db.Where("uuid = ?", videoUUID).First(&video)

	if result.Error != nil {
		return 0, result.Error
	}
	return video.ID, nil
}
