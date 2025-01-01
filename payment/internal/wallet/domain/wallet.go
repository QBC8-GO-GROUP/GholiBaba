package domain

import (
	"github.com/google/uuid"
	"time"
)

type WalletID uint

type Wallet struct {
	Id        WalletID
	UserId    string
	Code      uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
	Money     float64
}
