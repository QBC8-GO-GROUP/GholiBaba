package types

import (
	"gorm.io/gorm"
	"time"
)

type Wallet struct {
	gorm.Model
	Id        uint //WalletID
	UserId    string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
	Money     float64
}
