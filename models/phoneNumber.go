package models

import "gorm.io/gorm"

type PhoneNumber struct {
	gorm.Model
	Number string `json:"number" gorm:"not null;size:11;unique"`
	UserID uint   `json:"user_id"`
}
