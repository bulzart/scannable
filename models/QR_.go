package models

import "gorm.io/gorm"

type QR_ struct {
	gorm.Model
	UserID uint   `gorm:"not null" json:"user_id"`
	User   User   `gorm:"references:ID"`
	Url    string `gorm:"index,unique,not null"`
}
