package domain

import (
	"time"
)

type (
	TicketID         uint
	UserID           uint
	TravelID         uint
	TicketStatusType uint8
)

const (
	TicketStatusUnknown TicketStatusType = iota
	TicketStatusPending
	TicketStatusPaid
	TicketStatusCancelled
)

func (ts TicketStatusType) String() string {
	values := [...]string{
		"Unknown",
		"Pending",
		"Paid",
		"Cancelled",
	}
	return values[ts]
}

type Ticket struct {
	ID        TicketID         `json:"id"`
	UserID    UserID           `json:"user_id"`
	TravelID  TravelID         `json:"travel_id"`
	Status    TicketStatusType `json:"status"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
