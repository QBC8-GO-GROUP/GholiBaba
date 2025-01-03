package port

import (
	"context"
	"github.com/QBC8-GO-GROUP/GholiBaba/payment/internal/cards/domain"
)

type Service interface {
	CreateCard(ctx context.Context, card domain.Card) error
	UpdateCard(ctx context.Context, card domain.Card) error
	FindCardWithUserId(ctx context.Context, userId string) ([]domain.Card, error)
	DeleteCardWithUserId(ctx context.Context, userId string) error
	DeleteCardWithId(ctx context.Context, id int64) error
}
