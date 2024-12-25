package port

import (
	"context"

	"github.com/QBC8-GO-GROUP/GholiBaba/internal/travel/domain"
)

type Service interface {
	CreateTravel(ctx context.Context, company domain.Travel) (domain.TravelID, error)
	GetTravelByID(ctx context.Context, companyID domain.TravelID) (*domain.Travel, error)
	UpdateTravel(ctx context.Context, company domain.Travel) error
	DeleteByID(ctx context.Context, companyID domain.TravelID) error
}
