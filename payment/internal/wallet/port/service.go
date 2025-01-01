package port

import (
	"context"
	"github.com/QBC8-GO-GROUP/GholiBaba/payment/internal/wallet/domain"
)

type Service interface {
	CreateWallet(ctx context.Context, wallet domain.Wallet) error
	UpdateWallet(ctx context.Context, wallet domain.Wallet) error
	FindUserWallet(ctx context.Context, userId string) (domain.Wallet, error)
	DeleteUserWallet(ctx context.Context, walletId domain.WalletID) error
}
