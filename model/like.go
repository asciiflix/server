package model

import "gorm.io/gorm"

type Like struct {
	gorm.Model
	VideoID uint
	UserID  uint
}
