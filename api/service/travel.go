package service

import (
	"context"

	"github.com/QBC8-GO-GROUP/GholiBaba/internal/travel/domain"
	"github.com/QBC8-GO-GROUP/GholiBaba/internal/travel/port"
)

type TravelService struct {
	svc port.Service
}

func NewTravelService(svc port.Service) *TravelService {
	return &TravelService{svc: svc}
}

func (rs *TravelService) CreateTravel(ctx context.Context, travel domain.Travel) (domain.TravelID, error) {
	return rs.svc.CreateTravel(ctx, travel)
}

func (rs *TravelService) UpdateTravel(ctx context.Context, travel domain.Travel) error {
	return rs.svc.UpdateTravel(ctx, travel)
}

func (rs *TravelService) GetTravelByID(ctx context.Context, travelID domain.TravelID) (*domain.Travel, error) {
	return rs.svc.GetTravelByID(ctx, travelID)
}

func (rs *TravelService) GetTravels(ctx context.Context, companyID domain.OwnerID, page, pageSize int) ([]*domain.Travel, error) {
	return rs.svc.GetTravels(ctx, companyID, page, pageSize)
}

func (rs *TravelService) DeleteTravel(ctx context.Context, travelID domain.TravelID) error {
	return rs.svc.DeleteTravel(ctx, travelID)
}

func (rs *TravelService) BookTravel(ctx context.Context, travelID domain.TravelID) error {
	return rs.svc.BookTravel(ctx, travelID)
}

func (rs *TravelService) CancelBooking(ctx context.Context, travelID domain.TravelID) error {
	return rs.svc.CancelBooking(ctx, travelID)
}

func (rs *TravelService) ApproveTravel(ctx context.Context, travelID domain.TravelID) error {
	return rs.svc.ApproveTravel(ctx, travelID)
}

// TODO: CancelTravel should be a grpc service to be called from timer or something
// TODO: FinishTravel should be a grpc service to be called from timer or something
