package port

import (
	"context"
	"github.com/QBC8-GO-GROUP/GholiBaba/payment/internal/history/domain"
)

type Repo interface {
	Create(ctx context.Context, history domain.History) (domain.HistoryId, error)
	Update(ctx context.Context, history domain.History) error
	FindWithId(ctx context.Context, id domain.HistoryId) ([]domain.History, error)
	FindWithUserId(ctx context.Context, userId string) ([]domain.History, error)
	Delete(ctx context.Context, historyId domain.HistoryId) error
}
