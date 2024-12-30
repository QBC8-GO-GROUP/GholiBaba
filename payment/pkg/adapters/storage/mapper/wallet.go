package mapper

import (
	"github.com/QBC8-GO-GROUP/GholiBaba/payment/internal/wallet/domain"
	"github.com/QBC8-GO-GROUP/GholiBaba/payment/pkg/adapters/storage/types"
	"github.com/google/uuid"
)

func WalletDomainToStorage(domainWallet domain.Wallet) types.Wallet {
	return types.Wallet{
		Id:        uint(domainWallet.Id),
		UserId:    domainWallet.UserId,
		Code:      domainWallet.Code.String(),
		CreatedAt: domainWallet.CreatedAt,
		UpdatedAt: domainWallet.UpdatedAt,
		DeletedAt: domainWallet.DeletedAt,
		Money:     domainWallet.Money,
	}
}

func WalletStorageToDomain(storageWallet types.Wallet) (domain.Wallet, error) {
	code, err := uuid.Parse(storageWallet.Code)
	if err != nil {
		return domain.Wallet{}, err
	}
	return domain.Wallet{
		Id:        domain.WalletID(storageWallet.Id),
		UserId:    storageWallet.UserId,
		Code:      code,
		CreatedAt: storageWallet.CreatedAt,
		UpdatedAt: storageWallet.UpdatedAt,
		DeletedAt: storageWallet.DeletedAt,
		Money:     storageWallet.Money,
	}, nil
}
