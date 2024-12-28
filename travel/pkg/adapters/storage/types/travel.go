package types

import (
	"time"

	"gorm.io/gorm"
)

type Travel struct {
	gorm.Model
	Owner       uint      `gorm:"column:company_id;not null"`
	Type        uint8     `gorm:"column:type;not null"`
	Source      string    `gorm:"column:source;not null"`
	Destination string    `gorm:"column:destination;not null"`
	StartTime   time.Time `gorm:"column:start_time;not null"`
	EndTime     time.Time `gorm:"column:end_time;not null"`
	Price       float64   `gorm:"column:price;not null"`
	Seats       int       `gorm:"column:seats;not null"`
	Available   int       `gorm:"column:available;not null"`
	Approved    bool      `gorm:"column:approved;not null"`
	Vehicle     uint      `gorm:"column:vehicle_id;not null"`
}
