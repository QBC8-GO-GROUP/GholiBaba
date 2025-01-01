package port

import (
	"context"
	"github.com/QBC8-GO-GROUP/GholiBaba/payment/internal/wallet/domain"
)

type Repo interface {
	Create(ctx context.Context, wallet domain.Wallet) error
	Update(ctx context.Context, wallet domain.Wallet) error
	FindWithUserId(ctx context.Context, userId string) (domain.Wallet, error)
	DeleteWithUserId(ctx context.Context, walletId domain.WalletID) error
}
