package types

import (
	"gorm.io/gorm"
	"time"
)

type Wallet struct {
	gorm.Model
	Id        uint      `json:"id"`
	UserId    string    `json:"user_id"`
	Code      string    `json:"code"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
	Money     float64   `json:"money"`
}
