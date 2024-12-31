package domain

import (
	"github.com/QBC8-GO-GROUP/GholiBaba/payment/internal/wallet/domain"
	"github.com/google/uuid"
	"time"
)

type (
	HistoryId uint
)

type History struct {
	Id          HistoryId
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
	Code        uuid.UUID
	IsApproved  bool
	Price       float64
	Source      domain.WalletID
	Destination domain.WalletID
	Title       string
	Description string
}
