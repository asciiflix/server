package model

import "github.com/gofrs/uuid"

type Video struct {
	ID          uuid.UUID `gorm:"primary_key; unique; type:uuid; column:id;"`
	Title       string
	Description string
}
