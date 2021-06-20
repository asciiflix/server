package model

type Like struct {
	ID      uint `gorm:"primaryKey"`
	VideoID uint
	UserID  uint
}
