package types

import (
	"gorm.io/gorm"
	"time"
)

type Card struct {
	gorm.Model
	Id        uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
	Numbers   string
	WalletId  uint //walletDomain.WalletID
}
