package port

import (
	"context"

	"github.com/QBC8-GO-GROUP/GholiBaba/internal/travel/domain"
)

type Service interface {
	CreateTravel(ctx context.Context, travel domain.Travel) (domain.TravelID, error)
	UpdateTravel(ctx context.Context, travel domain.Travel) error
	GetTravelByID(ctx context.Context, travelID domain.TravelID) (*domain.Travel, error)
	GetTravels(ctx context.Context, companyID domain.OwnerID, page, pageSize int) ([]*domain.Travel, error)
	DeleteTravel(ctx context.Context, travelID domain.TravelID) error
	BookTravel(ctx context.Context, travelID domain.TravelID) error
	CancelBooking(ctx context.Context, travelID domain.TravelID) error
	CancelTravel(ctx context.Context, travelID domain.TravelID) error
	ApproveTravel(ctx context.Context, travelID domain.TravelID) error
	FinishTravel(ctx context.Context, travelID domain.TravelID) error
}
