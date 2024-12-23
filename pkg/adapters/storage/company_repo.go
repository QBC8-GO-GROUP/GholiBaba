package storage

import (
	"context"
	"errors"
	"time"

	"github.com/QBC8-GO-GROUP/GholiBaba/internal/company/domain"
	"github.com/QBC8-GO-GROUP/GholiBaba/internal/company/port"
	"github.com/QBC8-GO-GROUP/GholiBaba/pkg/adapters/storage/mapper"
	"github.com/QBC8-GO-GROUP/GholiBaba/pkg/adapters/storage/types"
	"github.com/QBC8-GO-GROUP/GholiBaba/pkg/logger"
	"gorm.io/gorm"
)

type companyRepo struct {
	db *gorm.DB
}

func NewCompanyRepo(db *gorm.DB) port.Repo {
	return &companyRepo{db}

}

func (r *companyRepo) Create(ctx context.Context, companyDomain domain.Company) (domain.CompanyID, error) {
	company := mapper.CompanyDomain2Storage(companyDomain)
	return domain.CompanyID(company.ID), r.db.Table("companys").WithContext(ctx).Create(company).Error
}

func (r *companyRepo) GetByID(ctx context.Context, companyID domain.CompanyID) (*domain.Company, error) {
	var company types.Company
	err := r.db.Debug().Table("companys").
		Where("id = ?", companyID).WithContext(ctx).
		First(&company).Error

	if err != nil {
		return nil, err
	}

	if company.ID == 0 {
		return nil, nil
	}

	return mapper.CompanyStorage2Domain(company), nil
}
func (r *companyRepo) GetByEmail(ctx context.Context, email domain.Email) (*domain.Company, error) {
	var company types.Company
	err := r.db.Table("companys").
		Where("email = ?", email).
		First(&company).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	if company.ID == 0 {
		return nil, nil
	}

	return mapper.CompanyStorage2Domain(company), nil
}

func (r *companyRepo) GetByFilter(ctx context.Context, filter *domain.CompanyFilter) (*domain.Company, error) {
	var company types.Company

	q := r.db.Table("companys").Debug().WithContext(ctx)

	if filter.ID > 0 {
		q = q.Where("id = ?", filter.ID)
	}

	if len(filter.Phone) > 0 {
		q = q.Where("phone = ?", filter.Phone)
	}

	err := q.First(&company).Error

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	if company.ID == 0 {
		return nil, nil
	}

	return mapper.CompanyStorage2Domain(company), nil
}

func (r *companyRepo) UpdateCompany(ctx context.Context, company domain.Company) error {
	var preUpdateCompany types.Company
	err := r.db.Model(&types.Company{}).Where("id = ?", company.ID).First((&preUpdateCompany)).Error
	if err != nil {
		logger.Error(err.Error(), nil)
		return err
	}
	currentTime := time.Now()
	if currentTime.Sub(preUpdateCompany.CreatedAt) > 24*time.Hour {
		return errors.New("can not update company due to limitation of update time")
	}
	updates := make(map[string]interface{})
	if company.FirstName != "" {
		updates["first_name"] = company.FirstName
	}

	if company.FirstName != "" {
		updates["last_name"] = company.LastName
	}

	if company.Phone != "" {
		updates["phone"] = company.Phone
	}

	if company.Email != "" {
		updates["email"] = company.Email
	}

	if company.NationalCode != "" {
		updates["national_code"] = company.NationalCode
	}

	if company.BirthDate != preUpdateCompany.BirthDate {
		updates["birth_date"] = company.BirthDate
	}

	if company.City != "" {
		updates["city"] = company.City
	}

	if company.Gender != preUpdateCompany.Gender {
		updates["gender"] = company.Gender
	}

	if company.SurveyLimitNumber != preUpdateCompany.SurveyLimitNumber {
		updates["survey_limit_number"] = company.SurveyLimitNumber
	}

	tx := r.db.Begin()
	if tx.Error != nil {
		logger.Error(tx.Error.Error(), nil)
		return tx.Error
	}

	// Update the company record
	if err := tx.Model(&types.Company{}).Where("id = ?", company.ID).Updates(updates).Error; err != nil {
		logger.Error(err.Error(), nil)
		tx.Rollback()
		return err
	}

	// Commit the transaction
	return tx.Commit().Error
}

func (r *companyRepo) DeleteByID(ctx context.Context, companyID domain.CompanyID) error {
	return r.db.Where("id = ?", companyID).Delete(&types.Company{}).Error
}
