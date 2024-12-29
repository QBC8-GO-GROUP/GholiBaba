package service

import (
	"context"

	"github.com/QBC8-GO-GROUP/GholiBaba/ticket/internal/ticket/domain"
	"github.com/QBC8-GO-GROUP/GholiBaba/ticket/internal/ticket/port"
)

type TicketService struct {
	svc port.Service
}

func NewTicketService(svc port.Service) *TicketService {
	return &TicketService{svc: svc}
}

func (rs *TicketService) BuyTicket(ctx context.Context, travelID domain.TravelID, userID domain.UserID) (domain.TicketID, error) {
	return rs.svc.BuyTicket(ctx, travelID, userID)
}

func (rs *TicketService) CancelTicket(ctx context.Context, ticketID domain.TicketID) error {
	return rs.svc.CancelTicket(ctx, ticketID)
}

func (rs *TicketService) GetUserTickets(ctx context.Context, userID domain.UserID) ([]*domain.Ticket, error) {
	return rs.svc.GetUserTickets(ctx, userID)
}
