package domain

import (
	walletDomain "github.com/QBC8-GO-GROUP/GholiBaba/payment/internal/wallet/domain"
	"time"
)

type Card struct {
	Id        uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
	Numbers   string
	WalletId  walletDomain.WalletID
}
