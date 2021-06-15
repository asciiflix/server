package database

import (
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
	result := global_db.Where("id = ?", videoId).First(&video)
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
func UpdateVideo(videoId string, video model.Video) error {
	//Get Video
	var videoToUpdate model.Video
	result := global_db.Where("id = ?", videoId).First(&videoToUpdate)
	if result.Error != nil {
		return result.Error
	}
	//Replaces non-zero fields
	result = global_db.Model(&videoToUpdate).Updates(video)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

//Delete video by id
func DeleteVideo(videoId string) error {
	result := global_db.Delete(&model.Video{}, videoId)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

//Testing for UUID-GEN
/*func CreateVideo() {
	var video model.VideoStats

	video.ID, _ = uuid.NewV4()
	video.Title = "Test Title"
	video.Description = "Test Desc"

	global_db.Create(&video)
}
*/
