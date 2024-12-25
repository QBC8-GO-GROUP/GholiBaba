package service

import (
	"context"

	"github.com/porseOnline/internal/user/domain"        // correct import after merging user logic
	userPort "github.com/porseOnline/internal/user/port" // correct import after merging user logic
)

type TravelService struct {
	svc                   userPort.TravelService
	authSecret            string
	expMin, refreshExpMin uint
}

func NewTravelService(svc userPort.TravelService, authSecret string, expMin, refreshExpMin uint) *TravelService {
	return &TravelService{svc: svc, authSecret: authSecret, expMin: expMin, refreshExpMin: refreshExpMin}
}

func (rs *TravelService) CreateTravel(ctx context.Context, role domain.Travel) (domain.TravelID, error) {
	return rs.svc.CreateTravel(ctx, role)
}

func (rs *TravelService) GetTravel(ctx context.Context, roleID domain.TravelID) (*domain.Travel, error) {
	return rs.svc.GetTravel(ctx, roleID)
}

func (rs *TravelService) UpdateTravel(ctx context.Context, role domain.Travel) error {
	return rs.svc.UpdateTravel(ctx, role)
}

func (rs *TravelService) DeleteTravel(ctx context.Context, roleID domain.TravelID) error {
	return rs.svc.DeleteTravel(ctx, roleID)
}
