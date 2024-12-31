package types

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type History struct {
	gorm.Model
	Id          uuid.UUID `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   time.Time `json:"deleted_at"`
	Code        uuid.UUID `json:"code"`
	IsApproved  bool      `json:"is_approved"`
	Price       float64   `json:"price"`
	Source      uint      `json:"source"`
	Destination uint      `json:"destination"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
}
