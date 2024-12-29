package port

import (
	"context"

	"github.com/QBC8-GO-GROUP/GholiBaba/ticket/internal/ticket/domain"
)

type Repo interface {
	Create(ctx context.Context, ticket domain.Ticket) (domain.TicketID, error)
	Delete(ctx context.Context, ticketID domain.TicketID) error
	GetTickets(ctx context.Context, userID domain.UserID) ([]*domain.Ticket, error)
}
