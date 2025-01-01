package history

import (
	"context"
	"errors"
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
	if err := validateHistory(history); err != nil {
		return 0, err
	}
	return s.repo.Create(ctx, history)
}

func (s *service) UpdateHistory(ctx context.Context, history domain.History) error {
	if err := validateHistory(history); err != nil {
		return err
	}
	return s.repo.Update(ctx, history)
}

func (s *service) FindHistoryWithId(ctx context.Context, id domain.HistoryId) ([]domain.History, error) {
	if id <= 0 {
		return nil, errors.New("invalid history ID")
	}
	return s.repo.FindWithId(ctx, id)
}

func (s *service) FindHistoryWithUserId(ctx context.Context, userId string) ([]domain.History, error) {
	if len(userId) == 0 {
		return nil, errors.New("invalid user ID")
	}
	return s.repo.FindWithUserId(ctx, userId)
}

func (s *service) DeleteHistory(ctx context.Context, historyId domain.HistoryId) error {
	if historyId <= 0 {
		return errors.New("invalid history ID")
	}
	return s.repo.Delete(ctx, historyId)
}

func validateHistory(history domain.History) error {
	if history.Price <= 0 {
		return errors.New("price must be greater than 0")
	}
	if history.Source <= 0 || history.Destination <= 0 {
		return errors.New("source and destination wallets must be valid")
	}
	if history.Title == "" {
		return errors.New("title cannot be empty")
	}
	return nil
}
