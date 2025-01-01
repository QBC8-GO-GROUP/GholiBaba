package port

import (
	"context"

	"github.com/QBC8-GO-GROUP/GholiBaba/internal/user/domain"
)

// input ports
type Repo interface {
	Create(ctx context.Context, user domain.User) (domain.UserID, error)
	GetByFilter(ctx context.Context, filter *domain.UserFilter) (*domain.User, error)
	GetById(ctx context.Context, id domain.UserID) (domain.User, error)
	Update(ctx context.Context, user domain.User) error
}
