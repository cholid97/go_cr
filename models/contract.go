package models

import "time"

type Contract struct {
	ID                  uint64    `gorm:"primary_key;auto_increment"`
	OTR                 int32     `json:"otr"`
	UserId              int32     `json:"user_id"`
	AdminFee            int32     `json:"admin_fee"`
	InstallmentAmount   int32     `json:"installment_amount"`
	InstallmentInterest int32     `json:"installment_interest"`
	Asset               string    `json:"assets"`
	CreatedAt           time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt           time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
}
