package types

import (
	"gorm.io/gorm"
	"time"
)

type Card struct {
	gorm.Model
	Id        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
	Numbers   string    `json:"numbers"`
	WalletId  uint      `json:"wallet_id"`
}
