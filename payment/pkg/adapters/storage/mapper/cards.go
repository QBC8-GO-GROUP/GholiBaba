package mapper

import (
	"github.com/QBC8-GO-GROUP/GholiBaba/payment/internal/cards/domain"
	walletDomain "github.com/QBC8-GO-GROUP/GholiBaba/payment/internal/wallet/domain"
	"github.com/QBC8-GO-GROUP/GholiBaba/payment/pkg/adapters/storage/types"
)

func CardsDomainToStorage(domainCard domain.Card) types.Card {
	return types.Card{
		Id:        domainCard.Id,
		Numbers:   domainCard.Numbers,
		WalletId:  uint(domainCard.WalletId),
		CreatedAt: domainCard.CreatedAt,
		UpdatedAt: domainCard.UpdatedAt,
		DeletedAt: domainCard.DeletedAt,
	}
}

func CardsStorageToDomain(storageCard types.Card) domain.Card {
	return domain.Card{
		Id:        storageCard.Id,
		Numbers:   storageCard.Numbers,
		WalletId:  walletDomain.WalletID(storageCard.WalletId),
		CreatedAt: storageCard.CreatedAt,
		UpdatedAt: storageCard.UpdatedAt,
		DeletedAt: storageCard.DeletedAt,
	}
}

func CardsStorageToDomainList(storageCards []types.Card) []domain.Card {
	domainCards := make([]domain.Card, len(storageCards))
	for i, card := range storageCards {
		domainCards[i] = CardsStorageToDomain(card)
	}
	return domainCards
}
