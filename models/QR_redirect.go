package models

import (
	"gorm.io/gorm"
	"time"
)

type QR_redirect struct {
	gorm.Model
	QrID      uint      `gorm:"not null" json:"qr_id"`
	QR        QR_       `gorm:"references:ID"`
	Url       string    `binding:"required"`
	StartDate time.Time `json:"start_date" binding:"required"`
	EndDate   time.Time `json:"end_date" binding:"required"`
}
