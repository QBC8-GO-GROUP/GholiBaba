package types

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type History struct {
	gorm.Model
	Id          uint //HistoryId
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
	Code        uuid.UUID
	IsApproved  bool
	Price       float64
	Source      uint //domain.WalletID
	Destination uint //domain.WalletID
	Title       string
	Description string
}
