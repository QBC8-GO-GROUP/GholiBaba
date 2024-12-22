package port

import "transportAgency/internal/agency/domain"

type AgencyRepository interface {
	Save(agency domain.TransportAgency) error
	FindByID(id string) (domain.TransportAgency, error)
}
