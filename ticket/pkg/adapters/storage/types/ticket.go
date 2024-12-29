package types

import (
	"gorm.io/gorm"
)

type Ticket struct {
	gorm.Model
	UserID   uint  `gorm:"column:user_id;not null"`
	TravelID uint  `gorm:"column:travel_id;not null"`
	Status   uint8 `gorm:"column:status;not null"`
}
