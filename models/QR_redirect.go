package models

import (
	"gorm.io/gorm"
	"time"
)

type QR_redirect struct {
	gorm.Model
	QRID      uint `gorm:"not null" json:"qr_id"`
	QR        QR_  `gorm:"references:ID"`
	Url       string
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
}
