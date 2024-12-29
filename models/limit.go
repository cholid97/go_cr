package models

import "time"

type Limit struct {
	ID           uint64    `gorm:"primary_key;auto_increment"`
	UserId       int32     `json:"user_id"`
	MonthlyLimit int32     `json:"month_limit"`
	MonthLimit   time.Time `json:"monthly_limit" gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	CreatedAt    time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt    time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
}
