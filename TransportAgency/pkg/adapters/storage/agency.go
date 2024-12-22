package storage

import (
	"errors"
	"transportAgency/internal/agency/domain"
)

type InMemoryAgencyRepo struct {
	agencies map[string]domain.TransportAgency
}

func NewInMemoryAgencyRepo() *InMemoryAgencyRepo {
	return &InMemoryAgencyRepo{
		agencies: make(map[string]domain.TransportAgency),
	}
}

func (r *InMemoryAgencyRepo) Save(agency domain.TransportAgency) error {
	if agency.Name == "" {
		return errors.New("agency name cannot be empty")
	}
	r.agencies[agency.ID] = agency
	return nil
}

func (r *InMemoryAgencyRepo) FindByID(id string) (domain.TransportAgency, error) {
	agency, exists := r.agencies[id]
	if !exists {
		return domain.TransportAgency{}, errors.New("agency not found")
	}
	return agency, nil
}
