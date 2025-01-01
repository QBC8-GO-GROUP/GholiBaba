package port

import (
	"context"
	"github.com/QBC8-GO-GROUP/GholiBaba/payment/internal/history/domain"
	"github.com/google/uuid"
)

type Repo interface {
	Create(ctx context.Context, history domain.History) (domain.HistoryId, error)
	Update(ctx context.Context, history domain.History) error
	FindWithId(ctx context.Context, id domain.HistoryId) ([]domain.History, error)
	FindWithCode(ctx context.Context, userId uuid.UUID) ([]domain.History, error)
	FindWithUserId(ctx context.Context, userId string) ([]domain.History, error)
	Delete(ctx context.Context, historyId domain.HistoryId) error
}
