package domain

import (
	walletDomain "github.com/QBC8-GO-GROUP/GholiBaba/payment/internal/wallet/domain"
	"time"
)

type Card struct {
	Id        uint                  `json:"id,omitempty"`
	CreatedAt time.Time             `json:"createdAt"`
	UpdatedAt time.Time             `json:"updatedAt"`
	DeletedAt time.Time             `json:"deletedAt"`
	Numbers   string                `json:"numbers,omitempty"`
	WalletId  walletDomain.WalletID `json:"walletId,omitempty"`
}
