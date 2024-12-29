package ticket

import (
	"context"
	"fmt"

	"github.com/QBC8-GO-GROUP/GholiBaba/ticket/internal/ticket/domain"
	"github.com/QBC8-GO-GROUP/GholiBaba/ticket/internal/ticket/port"

	"github.com/QBC8-GO-GROUP/GholiBaba/ticket/pkg/logger"
)

type service struct {
	repo port.Repo
}

func NewService(repo port.Repo) port.Service {
	return &service{
		repo: repo,
	}
}

func (s *service) BuyTicket(ctx context.Context, travelID domain.TravelID, userID domain.UserID) (domain.TicketID, error) {
	var newTicket = domain.Ticket{TravelID: travelID, UserID: userID, Status: domain.TicketStatusPending}
	ticketID, err := s.repo.Create(ctx, newTicket)
	if err != nil {
		logger.Error(fmt.Sprintf("error on buying new ticket %s", err.Error()), nil)
		return 0, err
	}

	return ticketID, nil
}

func (s *service) CancelTicket(ctx context.Context, ticketID domain.TicketID) error {
	err := s.repo.Delete(ctx, ticketID)
	if err != nil {
		logger.Error(fmt.Sprintf("error on canceling ticket %s", err.Error()), nil)
		return err
	}
	return nil
}

func (s *service) GetUserTickets(ctx context.Context, userID domain.UserID) ([]*domain.Ticket, error) {
	tickets, err := s.repo.GetTickets(ctx, userID)
	if err != nil {
		logger.Error(fmt.Sprintf("error on fetching tickets owned by a user %s", err.Error()), nil)
		return nil, err
	}

	return tickets, nil
}
