package travel

import (
	"context"
	"fmt"

	"github.com/QBC8-GO-GROUP/GholiBaba/internal/travel/domain"
	"github.com/QBC8-GO-GROUP/GholiBaba/internal/travel/port"

	"github.com/QBC8-GO-GROUP/GholiBaba/pkg/logger"
)

type service struct {
	repo port.Repo
}

func NewService(repo port.Repo) port.Service {
	return &service{
		repo: repo,
	}
}

func (s *service) CreateTravel(ctx context.Context, travel domain.Travel) (domain.TravelID, error) {
	travelID, err := s.repo.Create(ctx, travel)
	if err != nil {
		logger.Error(fmt.Sprintf("error on creating new travel %s", err.Error()), nil)
		return 0, err
	}

	return travelID, nil
}

func (s *service) UpdateTravel(ctx context.Context, travel domain.Travel) error {
	err := s.repo.Update(ctx, travel)
	if err != nil {
		logger.Error(fmt.Sprintf("error on updating travel %s", err.Error()), nil)
		return err
	}

	return nil
}

func (s *service) GetTravelByID(ctx context.Context, travelID domain.TravelID) (*domain.Travel, error) {
	travel, err := s.repo.Get(ctx, travelID)
	if err != nil {
		logger.Error(fmt.Sprintf("error on fetching travel by ID %s", err.Error()), nil)
		return nil, err
	}

	return travel, nil
}

func (s *service) GetTravels(ctx context.Context, companyID domain.OwnerID, page, pageSize int) ([]*domain.Travel, error) {
	travels, err := s.repo.GetAll(ctx, companyID, page, pageSize)
	if err != nil {
		logger.Error(fmt.Sprintf("error on fetching travels owned by a company %s", err.Error()), nil)
		return nil, err
	}

	return travels, nil
}

func (s *service) DeleteTravel(ctx context.Context, travelID domain.TravelID) error {
	err := s.repo.Delete(ctx, travelID)
	if err != nil {
		logger.Error(fmt.Sprintf("error on deleting travel %s", err.Error()), nil)
		return err
	}
	return nil
}

func (s *service) BookTravel(ctx context.Context, travelID domain.TravelID) error {
	err := s.repo.Book(ctx, travelID)
	if err != nil {
		logger.Error(fmt.Sprintf("error on booking travel %s", err.Error()), nil)
		return err
	}
	return nil
}

func (s *service) CancelBooking(ctx context.Context, travelID domain.TravelID) error {
	err := s.repo.Cancel(ctx, travelID)
	if err != nil {
		logger.Error(fmt.Sprintf("error on cancel travel book %s", err.Error()), nil)
		return err
	}
	return nil
}

func (s *service) CancelTravel(ctx context.Context, travelID domain.TravelID) error {
	err := s.repo.Delete(ctx, travelID)
	/*
		TODO: send notif to all ticket buyers that travel will be canceled
	*/
	if err != nil {
		logger.Error(fmt.Sprintf("error on canceling travel by ID %s", err.Error()), nil)
		return err
	}
	return nil
}

func (s *service) ApproveTravel(ctx context.Context, travelID domain.TravelID) error {
	err := s.repo.Approve(ctx, travelID)
	if err != nil {
		logger.Error(fmt.Sprintf("error on approving travel by ID %s", err.Error()), nil)
		return err
	}
	return nil
}

func (s *service) FinishTravel(ctx context.Context, travelID domain.TravelID) error {
	/*
		TODO: check with map service if travel finish is confirmed
		TODO: send finish signal to payments to share the money
	*/
	return nil
}
