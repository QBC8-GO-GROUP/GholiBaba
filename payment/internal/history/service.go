package history

import (
	"context"
	"github.com/QBC8-GO-GROUP/GholiBaba/payment/internal/history/domain"
	"github.com/QBC8-GO-GROUP/GholiBaba/payment/internal/history/port"
)

type service struct {
	repo port.Repo
}

func NewService(r port.Repo) port.Service {
	return &service{repo: r}
}

func (s *service) CreateHistory(ctx context.Context, history domain.History) (domain.HistoryId, error) {
	//TODO implement me
	panic("implement me")
}

func (s *service) UpdateHistory(ctx context.Context, history domain.History) error {
	//TODO implement me
	panic("implement me")
}

func (s *service) FindHistoryWithId(ctx context.Context, id domain.HistoryId) ([]domain.History, error) {
	//TODO implement me
	panic("implement me")
}

func (s *service) FindHistoryWithUserId(ctx context.Context, userId int64) ([]domain.History, error) {
	//TODO implement me
	panic("implement me")
}

func (s *service) DeleteHistory(ctx context.Context, historyId domain.HistoryId) error {
	//TODO implement me
	panic("implement me")
}
