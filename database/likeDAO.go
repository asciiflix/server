package database

import (
	"errors"

	"github.com/asciiflix/server/model"
	"github.com/asciiflix/server/utils"
	"gorm.io/gorm"
)

func CheckIfLiked(videoID uint, userID string) (bool, error) {
	result := global_db.Where("video_id = ?", videoID).Where("user_id = ?", userID).First(&model.Like{})
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return true, nil
		}
	}
	return false, result.Error

}

//Create video
func CreateLike(videoID uint, userID string) error {

	//Check if liked
	result := global_db.Where("video_id = ?", videoID).Where("user_id = ?", userID).First(&model.Like{})
	if result.Error != gorm.ErrRecordNotFound {
		return errors.New("already liked")
	}

	vidID := videoID

	useID, err := utils.ParseStringToUint(userID)
	if err != nil {
		return err
	}

	like := model.Like{
		VideoID: vidID,
		UserID:  useID,
	}

	//Create Like
	result = global_db.Create(&like)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func DeleteLike(videoID uint, userID string) error {
	result := global_db.Where("video_id = ?", videoID).Where("user_id = ?", userID).Delete(&model.Like{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}
