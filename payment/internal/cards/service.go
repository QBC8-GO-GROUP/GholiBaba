package cards

import (
	"context"
	"errors"
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
	if err := validateCard(card); err != nil {
		return err
	}
	return s.repo.Create(ctx, card)
}

func (s *service) UpdateCard(ctx context.Context, card domain.Card) error {
	if err := validateCard(card); err != nil {
		return err
	}
	return s.repo.Update(ctx, card)
}

func (s *service) FindCardWithUserId(ctx context.Context, userId string) ([]domain.Card, error) {
	if len(userId) == 0 {
		return nil, errors.New("invalid user ID")
	}
	return s.repo.FindWithUserId(ctx, userId)
}

func (s *service) DeleteCardWithUserId(ctx context.Context, userId string) error {
	if len(userId) == 0 {
		return errors.New("invalid user ID")
	}
	return s.repo.DeleteWithUserId(ctx, userId)
}

func (s *service) DeleteCardWithId(ctx context.Context, id int64) error {
	if id <= 0 {
		return errors.New("invalid card ID")
	}
	return s.repo.DeleteWithId(ctx, id)
}

func validateCard(card domain.Card) error {
	if len(card.Numbers) != 16 {
		return errors.New("card number must be 16 digits")
	}
	if card.WalletId <= 0 {
		return errors.New("invalid wallet ID")
	}
	return nil
}
