package service

import (
	"github.com/google/uuid"
	"transportAgency/internal/agency/domain"
	"transportAgency/internal/agency/port"
)

type CreateAgencyCommand struct {
	Name string
}

type CreateAgencyService struct {
	Repo port.AgencyRepository
}

func (s *CreateAgencyService) Execute(cmd CreateAgencyCommand) (string, error) {
	agency := domain.TransportAgency{
		ID:   uuid.NewString(),
		Name: cmd.Name,
	}

	err := s.Repo.Save(agency)
	return agency.ID, err
}
