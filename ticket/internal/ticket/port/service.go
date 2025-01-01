package port

import (
	"context"

	"github.com/QBC8-GO-GROUP/GholiBaba/ticket/internal/ticket/domain"
)

type Service interface {
	BuyTicket(ctx context.Context, travelID domain.TravelID, userID domain.UserID) (domain.TicketID, error)
	CancelTicket(ctx context.Context, ticketID domain.TicketID) error
	GetUserTickets(ctx context.Context, userID domain.UserID) ([]*domain.Ticket, error)
}
