package company

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/QBC8-GO-GROUP/GholiBaba/internal/company/domain"
	"github.com/QBC8-GO-GROUP/GholiBaba/internal/company/port"

	"github.com/QBC8-GO-GROUP/GholiBaba/pkg/logger"

	"golang.org/x/crypto/bcrypt"
)

var (
	ErrCompanyOnCreate           = errors.New("error on creating new company")
	ErrCompanyCreationValidation = errors.New("validation failed")
	ErrCompanyNotFound           = errors.New("company not found")
)

type service struct {
	repo port.Repo
}

func NewService(repo port.Repo) port.Service {
	return &service{
		repo: repo,
	}
}

func (s *service) CreateCompany(ctx context.Context, company domain.Company) (domain.CompanyID, error) {
	if err := company.Validate(); err != nil {
		return 0, fmt.Errorf("%w %w", ErrCompanyCreationValidation, err)
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(company.PasswordHash), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Error while hashing password : ", err.Error())
		return 0, ErrCompanyOnCreate
	}
	company.PasswordHash = string(hashedPassword)
	companyID, err := s.repo.Create(ctx, company)
	if err != nil {
		log.Println("error on creating new company : ", err.Error())
		return 0, ErrCompanyOnCreate
	}

	return companyID, nil
}

func (s *service) GetCompanyByID(ctx context.Context, companyID domain.CompanyID) (*domain.Company, error) {
	company, err := s.repo.GetByID(ctx, companyID)
	if err != nil {
		return nil, err
	}
	if company == nil || company.ID == 0 {
		return nil, ErrCompanyNotFound
	}

	return company, nil
}

func (s *service) GetCompanyByEmail(ctx context.Context, email domain.Email) (*domain.Company, error) {
	company, err := s.repo.GetByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	if company == nil || company.ID == 0 {
		return nil, ErrCompanyNotFound
	}

	return company, nil
}

func (s *service) UpdateCompany(ctx context.Context, company domain.Company) error {
	err := s.repo.UpdateCompany(ctx, company)
	if err != nil {
		logger.Error("error in update company", nil)
		return err
	}
	return nil
}

func (s *service) DeleteByID(ctx context.Context, companyID domain.CompanyID) error {
	err := s.repo.DeleteByID(ctx, companyID)
	if err != nil {
		logger.Error("can not delete company", nil)
		return err
	}
	return nil
}
