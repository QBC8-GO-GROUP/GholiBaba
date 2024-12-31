package port

import (
	"context"
	"github.com/QBC8-GO-GROUP/GholiBaba/payment/internal/history/domain"
)

type Service interface {
	CreateHistory(ctx context.Context, history domain.History) (domain.HistoryId, error)
	UpdateHistory(ctx context.Context, history domain.History) error
	FindHistoryWithId(ctx context.Context, id domain.HistoryId) ([]domain.History, error)
	FindHistoryWithUserId(ctx context.Context, userId string) ([]domain.History, error)
	DeleteHistory(ctx context.Context, historyId domain.HistoryId) error
}
