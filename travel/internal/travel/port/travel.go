package port

import (
	"context"

	"github.com/QBC8-GO-GROUP/GholiBaba/internal/travel/domain"
)

type Repo interface {
	Create(ctx context.Context, travel domain.Travel) (domain.TravelID, error)
	Update(ctx context.Context, travel domain.Travel) error
	Get(ctx context.Context, travelID domain.TravelID) (*domain.Travel, error)
	GetAll(ctx context.Context, companyID domain.OwnerID, page, pageSize int) ([]*domain.Travel, error)
	Delete(ctx context.Context, travelID domain.TravelID) error
	Book(ctx context.Context, travelID domain.TravelID) error
	Cancel(ctx context.Context, travelID domain.TravelID) error
	Approve(ctx context.Context, travelID domain.TravelID) error
}
