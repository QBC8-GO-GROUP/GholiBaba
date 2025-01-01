package port

import (
	"context"

	"github.com/QBC8-GO-GROUP/GholiBaba/internal/user/domain"
)

// export ports
type Service interface {
	CreateUser(ctx context.Context, user domain.User) (domain.UserID, error)
	GetUserByFilter(ctx context.Context, filter *domain.UserFilter) (*domain.User, error)
	GetUserById(ctx context.Context, id domain.UserID) (domain.User, error)
	UpdateUser(ctx context.Context, user domain.User) error
}
