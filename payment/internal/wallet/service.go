package wallet

import (
	"context"
	"errors"
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
	if err := validateWallet(wallet); err != nil {
		return err
	}
	return s.repo.Create(ctx, wallet)
}

func (s *service) UpdateWallet(ctx context.Context, wallet domain.Wallet) error {
	if err := validateWallet(wallet); err != nil {
		return err
	}
	return s.repo.Update(ctx, wallet)
}

func (s *service) FindUserWallet(ctx context.Context, userId string) (domain.Wallet, error) {
	if userId == "" {
		return domain.Wallet{}, errors.New("user ID cannot be empty")
	}
	return s.repo.FindWithUserId(ctx, userId)
}

func (s *service) DeleteUserWallet(ctx context.Context, walletId domain.WalletID) error {
	if walletId <= 0 {
		return errors.New("invalid wallet ID")
	}
	return s.repo.DeleteWithUserId(ctx, walletId)
}

func validateWallet(wallet domain.Wallet) error {
	if wallet.UserId == "" {
		return errors.New("user ID cannot be empty")
	}
	if wallet.Money < 0 {
		return errors.New("money cannot be negative")
	}
	return nil
}
