package cards

import (
	"context"
	"github.com/QBC8-GO-GROUP/GholiBaba/payment/internal/cards/domain"
	"github.com/QBC8-GO-GROUP/GholiBaba/payment/internal/cards/port"
)

type service struct {
	repo port.Repo
}

func NewService(repo port.Repo) port.Service {
	return &service{repo}
}

// validate is here

func (s *service) CreateCard(ctx context.Context, card domain.Card) error {
	return s.repo.Create(ctx, card)
}

func (s *service) UpdateCard(ctx context.Context, card domain.Card) error {
	return s.repo.Update(ctx, card)
}

func (s *service) FindCardWithUserId(ctx context.Context, userId int64) ([]domain.Card, error) {
	return s.repo.FindWithUserId(ctx, userId)
}

func (s *service) DeleteCardWithUserId(ctx context.Context, userId int64) error {
	return s.repo.DeleteWithUserId(ctx, userId)
}

func (s *service) DeleteCardWithId(ctx context.Context, id int64) error {
	return s.repo.DeleteWithId(ctx, id)
}
