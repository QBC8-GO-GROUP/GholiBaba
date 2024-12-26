package mapper

import (
	"github.com/QBC8-GO-GROUP/GholiBaba/internal/travel/domain"
	"github.com/QBC8-GO-GROUP/GholiBaba/pkg/adapters/storage/types"

	"gorm.io/gorm"
)

func TravelDomain2Storage(travelDomain domain.Travel) *types.Travel {
	return &types.Travel{
		Model: gorm.Model{
			ID:        uint(travelDomain.ID),
			CreatedAt: travelDomain.CreatedAt,
			DeletedAt: gorm.DeletedAt(ToNullTime(travelDomain.DeletedAt)),
			UpdatedAt: travelDomain.UpdatedAt,
		},
		Owner:       uint(travelDomain.Owner),
		Type:        travelDomain.Type,
		Source:      travelDomain.Source,
		Destination: travelDomain.Destination,
		StartTime:   travelDomain.StartTime,
		EndTime:     travelDomain.EndTime,
		Price:       travelDomain.Price,
		Seats:       travelDomain.Seats,
		Available:   travelDomain.Available,
		Approved:    travelDomain.Approved,
		Vehicle:     uint(travelDomain.Vehicle),
	}
}

func TravelStorage2Domain(travel types.Travel) *domain.Travel {
	return &domain.Travel{
		ID:          domain.TravelID(travel.ID),
		CreatedAt:   travel.CreatedAt,
		UpdatedAt:   travel.UpdatedAt,
		Owner:       domain.OwnerID(travel.Owner),
		Type:        travel.Type,
		Source:      travel.Source,
		Destination: travel.Destination,
		StartTime:   travel.StartTime,
		EndTime:     travel.EndTime,
		Price:       travel.Price,
		Seats:       travel.Seats,
		Available:   travel.Available,
		Approved:    travel.Approved,
		Vehicle:     domain.VehicleID(travel.Vehicle),
	}
}
