package wallet

import (
	"context"
	"github.com/QBC8-GO-GROUP/GholiBaba/payment/internal/wallet/domain"
	"github.com/QBC8-GO-GROUP/GholiBaba/payment/internal/wallet/port"
)

type service struct {
	repo port.Repo
}

func NewService(repo port.Repo) port.Service {
	return &service{repo}
}

func (s *service) CreateWallet(ctx context.Context, wallet domain.Wallet) error {
	//TODO implement me
	panic("implement me")
}

func (s *service) UpdateWallet(ctx context.Context, wallet domain.Wallet) error {
	//TODO implement me
	panic("implement me")
}

func (s *service) FindUserWallet(ctx context.Context, userId string) (domain.Wallet, error) {
	//TODO implement me
	panic("implement me")
}

func (s *service) DeleteUserWallet(ctx context.Context, walletId domain.WalletID) error {
	//TODO implement me
	panic("implement me")
}
