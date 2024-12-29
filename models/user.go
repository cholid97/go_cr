package models

import (
	"time"
)

type User struct {
	ID        uint      `gorm:"primary_key;auto_increment"`
	NIK       int       `json:"nik"`
	FullName  string    `json:"fullname"`
	LegalName string    `json:"legalname"`
	DOB       string    `json:"dob"`
	POB       time.Time `json:"pob" gorm:"type:date"`
	CreatedAt time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
}
