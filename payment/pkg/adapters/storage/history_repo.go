package storage

import (
	"context"
	"github.com/QBC8-GO-GROUP/GholiBaba/payment/internal/history/domain"
	"github.com/QBC8-GO-GROUP/GholiBaba/payment/internal/history/port"
	"github.com/QBC8-GO-GROUP/GholiBaba/payment/pkg/adapters/storage/mapper"
	"github.com/QBC8-GO-GROUP/GholiBaba/payment/pkg/adapters/storage/types"
	"gorm.io/gorm"
)

type historyRepo struct {
	db *gorm.DB
}

func NewHistoryRepo(db *gorm.DB) port.Repo {
	return &historyRepo{db}
}

func (h *historyRepo) Create(ctx context.Context, history domain.History) (domain.HistoryId, error) {
	storageHistory := mapper.HistoryDomainToStorage(history)
	err := h.db.WithContext(ctx).Create(&storageHistory).Error
	return domain.HistoryId(storageHistory.Id), err
}

func (h *historyRepo) Update(ctx context.Context, history domain.History) error {
	storageHistory := mapper.HistoryDomainToStorage(history)
	return h.db.WithContext(ctx).
		Model(&types.History{}).
		Where("id = ?", storageHistory.Id).
		Updates(storageHistory).Error
}

func (h *historyRepo) FindWithId(ctx context.Context, id domain.HistoryId) ([]domain.History, error) {
	var storageHistories []types.History
	err := h.db.WithContext(ctx).
		Where("id = ?", uint(id)).
		Find(&storageHistories).Error
	if err != nil {
		return nil, err
	}
	return mapper.HistoryStorageToDomainList(storageHistories), nil
}

func (h *historyRepo) FindWithUserId(ctx context.Context, userId int64) ([]domain.History, error) {
	var storageHistories []types.History
	err := h.db.WithContext(ctx).
		Where("source = ? OR destination = ?", userId, userId).
		Find(&storageHistories).Error
	if err != nil {
		return nil, err
	}
	return mapper.HistoryStorageToDomainList(storageHistories), nil
}

func (h *historyRepo) Delete(ctx context.Context, historyId domain.HistoryId) error {
	return h.db.WithContext(ctx).
		Where("id = ?", uint(historyId)).
		Delete(&types.History{}).Error
}
