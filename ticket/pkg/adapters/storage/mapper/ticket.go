package mapper

import (
	"github.com/QBC8-GO-GROUP/GholiBaba/ticket/internal/ticket/domain"
	"github.com/QBC8-GO-GROUP/GholiBaba/ticket/pkg/adapters/storage/types"

	"gorm.io/gorm"
)

func TicketDomain2Storage(ticketDomain domain.Ticket) *types.Ticket {
	return &types.Ticket{
		Model: gorm.Model{
			ID:        uint(ticketDomain.ID),
			CreatedAt: ticketDomain.CreatedAt,
			DeletedAt: gorm.DeletedAt(ToNullTime(ticketDomain.DeletedAt)),
			UpdatedAt: ticketDomain.UpdatedAt,
		},
		UserID:   uint(ticketDomain.UserID),
		TravelID: uint(ticketDomain.TravelID),
		Status:   uint8(ticketDomain.Status),
	}
}

func TicketStorage2Domain(ticket types.Ticket) *domain.Ticket {
	return &domain.Ticket{
		ID:        domain.TicketID(ticket.ID),
		CreatedAt: ticket.CreatedAt,
		UpdatedAt: ticket.UpdatedAt,
		UserID:    domain.UserID(ticket.UserID),
		TravelID:  domain.TravelID(ticket.TravelID),
		Status:    domain.TicketStatusType(ticket.Status),
	}
}
