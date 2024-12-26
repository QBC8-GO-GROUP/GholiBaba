package types

import (
	"time"

	"gorm.io/gorm"
)

type Travel struct {
	gorm.Model
	Owner       uint      `gorm:"column:company_id"`
	Type        string    `gorm:"column:type"`
	Source      string    `gorm:"column:source"`
	Destination string    `gorm:"column:destination"`
	StartTime   time.Time `gorm:"column:start_time"`
	EndTime     time.Time `gorm:"column:end_time"`
	Price       float64   `gorm:"column:price"`
	Seats       int       `gorm:"column:seats"`
	Available   int       `gorm:"column:available"`
	Approved    bool      `gorm:"column:approved"`
	Vehicle     uint      `gorm:"column:vehicle_id"`
}
