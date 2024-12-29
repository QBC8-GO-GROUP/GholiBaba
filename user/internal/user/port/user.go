package port

import (
	"context"

	"github.com/QBC8-GO-GROUP/GholiBaba/internal/user/domain"
)

// input ports
type Repo interface {
	Create(ctx context.Context, user domain.User) (domain.UserID, error)
	GetByFilter(ctx context.Context, filter *domain.UserFilter) (*domain.User, error)
}
