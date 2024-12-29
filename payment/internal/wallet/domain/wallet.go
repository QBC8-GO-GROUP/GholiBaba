package domain

import "time"

type WalletID uint

type Wallet struct {
	Id        WalletID
	UserId    string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
	Money     float64
}
