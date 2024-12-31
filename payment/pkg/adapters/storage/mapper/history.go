package mapper

import (
	"github.com/QBC8-GO-GROUP/GholiBaba/payment/internal/history/domain"
	walletDomain "github.com/QBC8-GO-GROUP/GholiBaba/payment/internal/wallet/domain"
	"github.com/QBC8-GO-GROUP/GholiBaba/payment/pkg/adapters/storage/types"
	"github.com/google/uuid"
)

func HistoryDomainToStorage(domainHistory domain.History) types.History {
	return types.History{
		Id:          uuid.UUID(domainHistory.Id),
		CreatedAt:   domainHistory.CreatedAt,
		UpdatedAt:   domainHistory.UpdatedAt,
		DeletedAt:   domainHistory.DeletedAt,
		Code:        domainHistory.Code,
		IsApproved:  domainHistory.IsApproved,
		Price:       domainHistory.Price,
		Source:      uint(domainHistory.Source),
		Destination: uint(domainHistory.Destination),
		Title:       domainHistory.Title,
		Description: domainHistory.Description,
	}
}

func HistoryStorageToDomain(storageHistory types.History) domain.History {
	return domain.History{
		Id:          domain.HistoryId(storageHistory.Id),
		CreatedAt:   storageHistory.CreatedAt,
		UpdatedAt:   storageHistory.UpdatedAt,
		DeletedAt:   storageHistory.DeletedAt,
		Code:        storageHistory.Code,
		IsApproved:  storageHistory.IsApproved,
		Price:       storageHistory.Price,
		Source:      walletDomain.WalletID(storageHistory.Source),
		Destination: walletDomain.WalletID(storageHistory.Destination),
		Title:       storageHistory.Title,
		Description: storageHistory.Description,
	}
}

func HistoryStorageToDomainList(storageHistories []types.History) []domain.History {
	domainHistories := make([]domain.History, len(storageHistories))
	for i, history := range storageHistories {
		domainHistories[i] = HistoryStorageToDomain(history)
	}
	return domainHistories
}
