package domain

import (
	"time"
)

type (
	TravelID   uint
	OwnerID    uint
	VehicleID  uint
	TravelType uint8
)

const (
	TravelTypeUnknown TravelType = iota
	TravelTypeBus
	TravelTypeTrain
	TravelTypeFlight
	TravelTypeShip
)

func (tt TravelType) String() string {
	values := [...]string{
		"Unknown",
		"Bus",
		"Train",
		"Flight",
		"Ship",
	}
	return values[tt]
}

type Travel struct {
	ID          TravelID
	CreatedAt   time.Time
	DeletedAt   time.Time
	UpdatedAt   time.Time
	Owner       OwnerID    `json:"company_id"`
	Type        TravelType `json:"type"`
	Source      string     `json:"source"`
	Destination string     `json:"destination"`
	StartTime   time.Time  `json:"start_time"`
	EndTime     time.Time  `json:"end_time"`
	Price       float64    `json:"price"`
	Seats       int        `json:"seats"`
	Available   int        `json:"available"`
	Approved    bool       `json:"approved"`
	Vehicle     VehicleID  `json:"vehicle_id"`
}
