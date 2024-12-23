package port

import (
	"context"

	"github.com/QBC8-GO-GROUP/GholiBaba/internal/company/domain"
)

type Repo interface {
	Create(ctx context.Context, company domain.Company) (domain.CompanyID, error)
	GetByID(ctx context.Context, companyID domain.CompanyID) (*domain.Company, error)
	GetByEmail(ctx context.Context, email domain.Email) (*domain.Company, error)

	UpdateCompany(ctx context.Context, company domain.Company) error
	DeleteByID(ctx context.Context, companyID domain.CompanyID) error

	GetByFilter(ctx context.Context, filter *domain.CompanyFilter) (*domain.Company, error)
}
