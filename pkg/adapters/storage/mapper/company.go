package mapper

import (
	"github.com/babyhando/order-service/internal/company/domain"
	"github.com/babyhando/order-service/pkg/adapters/storage/types"

	"gorm.io/gorm"
)

func CompanyDomain2Storage(companyDomain domain.Company) *types.Company {
	return &types.Company{
		Model: gorm.Model{
			ID:        uint(companyDomain.ID),
			CreatedAt: companyDomain.CreatedAt,
			DeletedAt: gorm.DeletedAt(ToNullTime(companyDomain.DeletedAt)),
		},
		FirstName: companyDomain.FirstName,
		LastName:  companyDomain.LastName,
		Phone:     string(companyDomain.Phone),
		Password:  companyDomain.Password,
	}
}

func CompanyStorage2Domain(company types.Company) *domain.Company {
	return &domain.Company{
		ID:        domain.CompanyID(company.ID),
		CreatedAt: company.CreatedAt,
		DeletedAt: company.DeletedAt.Time,
		FirstName: company.FirstName,
		LastName:  company.LastName,
		Phone:     domain.Phone(company.Phone),
		Password:  company.Password,
	}
}
