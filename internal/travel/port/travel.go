package port

import (
	"context"

	"github.com/QBC8-GO-GROUP/GholiBaba/internal/travel/domain"
)

type Repo interface {
	Create(ctx context.Context, role domain.Travel) (domain.TravelID, error)
	Get(ctx context.Context, roleID domain.TravelID) (*domain.Travel, error)
	Update(ctx context.Context, role domain.Travel) error
	Delete(ctx context.Context, roleID domain.TravelID) error
}
