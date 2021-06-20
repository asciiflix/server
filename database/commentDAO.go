package database

import "github.com/asciiflix/server/model"

func CreateComment(comment model.Comment) error {

	result := global_db.Create(&comment)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetComments(videoID string) ([]model.Comment, error) {
	video, err := GetVideo(videoID)
	if err != nil {
		return nil, err
	}
	return video.Comments, nil
}
