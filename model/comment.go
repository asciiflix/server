package model

import "time"

type Comment struct {
	ID        uint `gorm:"primaryKey"`
	UserID    uint
	VideoID   uint
	CreatedAt time.Time
	Content   string `gorm:"not null"`
}
