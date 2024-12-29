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

func (s *service) CreateCard(ctx context.Context, card domain.Card) error {
	//TODO implement me
	panic("implement me")
}

func (s *service) UpdateCard(ctx context.Context, card domain.Card) error {
	//TODO implement me
	panic("implement me")
}

func (s *service) FindCardWithUserId(ctx context.Context, userId int64) ([]domain.Card, error) {
	//TODO implement me
	panic("implement me")
}

func (s *service) DeleteCardWithUserId(ctx context.Context, userId int64) error {
	//TODO implement me
	panic("implement me")
}

func (s *service) DeleteCardWithId(ctx context.Context, id int64) error {
	//TODO implement me
	panic("implement me")
}
