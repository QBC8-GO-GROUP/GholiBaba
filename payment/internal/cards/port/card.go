package port

import (
	"context"
	"github.com/QBC8-GO-GROUP/GholiBaba/payment/internal/cards/domain"
)

type Repo interface {
	Create(ctx context.Context, card domain.Card) error
	Update(ctx context.Context, card domain.Card) error
	FindWithUserId(ctx context.Context, userId string) ([]domain.Card, error)
	DeleteWithUserId(ctx context.Context, userId string) error
	DeleteWithId(ctx context.Context, id int64) error
}
