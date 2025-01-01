package storage

import (
	"context"
	"github.com/QBC8-GO-GROUP/GholiBaba/payment/internal/wallet/domain"
	"github.com/QBC8-GO-GROUP/GholiBaba/payment/internal/wallet/port"
	"github.com/QBC8-GO-GROUP/GholiBaba/payment/pkg/adapters/storage/mapper"
	"github.com/QBC8-GO-GROUP/GholiBaba/payment/pkg/adapters/storage/types"
	"gorm.io/gorm"
)

type walletRepo struct {
	db *gorm.DB
}

func NewWalletRepo(db *gorm.DB) port.Repo {
	return &walletRepo{db}
}

func (w *walletRepo) Create(ctx context.Context, wallet domain.Wallet) error {
	storageWallet := mapper.WalletDomainToStorage(wallet)
	return w.db.WithContext(ctx).Create(&storageWallet).Error
}

func (w *walletRepo) Update(ctx context.Context, wallet domain.Wallet) error {
	storageWallet := mapper.WalletDomainToStorage(wallet)
	return w.db.WithContext(ctx).
		Model(&types.Wallet{}).
		Where("id = ?", storageWallet.Id).
		Updates(storageWallet).Error
}

func (w *walletRepo) FindWithUserId(ctx context.Context, userId string) (domain.Wallet, error) {
	var storageWallet types.Wallet
	err := w.db.WithContext(ctx).
		Where("user_id = ?", userId).
		First(&storageWallet).Error
	if err != nil {
		return domain.Wallet{}, err
	}
	return mapper.WalletStorageToDomain(storageWallet)
}

func (w *walletRepo) DeleteWithUserId(ctx context.Context, walletId domain.WalletID) error {
	return w.db.WithContext(ctx).
		Where("id = ?", uint(walletId)).
		Delete(&types.Wallet{}).Error
}
