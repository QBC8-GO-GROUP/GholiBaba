package storage

import (
	"context"
	"github.com/QBC8-GO-GROUP/GholiBaba/payment/internal/cards/domain"
	"github.com/QBC8-GO-GROUP/GholiBaba/payment/internal/cards/port"
	"github.com/QBC8-GO-GROUP/GholiBaba/payment/pkg/adapters/storage/mapper"
	"github.com/QBC8-GO-GROUP/GholiBaba/payment/pkg/adapters/storage/types"
	"gorm.io/gorm"
)

type cardsRepo struct {
	db *gorm.DB
}

func NewCardsRepo(db *gorm.DB) port.Repo {
	return &cardsRepo{db}
}

func (c *cardsRepo) Create(ctx context.Context, card domain.Card) error {
	storageCard := mapper.CardsDomainToStorage(card)
	return c.db.WithContext(ctx).Create(&storageCard).Error
}

func (c *cardsRepo) Update(ctx context.Context, card domain.Card) error {
	storageCard := mapper.CardsDomainToStorage(card)
	return c.db.WithContext(ctx).Model(&types.Card{}).Where("id = ?", card.Id).Updates(storageCard).Error
}

func (c *cardsRepo) FindWithUserId(ctx context.Context, userId string) ([]domain.Card, error) {
	var storageCards []types.Card
	err := c.db.WithContext(ctx).Where("wallet_id = ?", userId).Find(&storageCards).Error
	if err != nil {
		return nil, err
	}
	return mapper.CardsStorageToDomainList(storageCards), nil
}

func (c *cardsRepo) DeleteWithUserId(ctx context.Context, userId string) error {
	return c.db.WithContext(ctx).Where("wallet_id = ?", userId).Delete(&types.Card{}).Error
}

func (c *cardsRepo) DeleteWithId(ctx context.Context, id int64) error {
	return c.db.WithContext(ctx).Where("id = ?", id).Delete(&types.Card{}).Error
}
