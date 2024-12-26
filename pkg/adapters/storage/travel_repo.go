package storage

import (
	"context"
	"errors"
	"time"

	"github.com/QBC8-GO-GROUP/GholiBaba/internal/travel/domain"
	"github.com/QBC8-GO-GROUP/GholiBaba/internal/travel/port"
	"github.com/QBC8-GO-GROUP/GholiBaba/pkg/adapters/storage/mapper"
	"github.com/QBC8-GO-GROUP/GholiBaba/pkg/adapters/storage/types"
	"github.com/QBC8-GO-GROUP/GholiBaba/pkg/logger"
	"gorm.io/gorm"
)

type travelRepo struct {
	db *gorm.DB
}

func NewTravelRepo(db *gorm.DB) port.Repo {
	return &travelRepo{db}

}

func (tr *travelRepo) Create(ctx context.Context, travelDomain domain.Travel) (domain.TravelID, error) {
	if travelDomain.Type != "" {
		logger.Error("travel type cannot be empty", nil)
		return 0, errors.New("travel type cannot be empty")
	}
	if travelDomain.Source != "" {
		logger.Error("travel source cannot be empty", nil)
		return 0, errors.New("travel source cannot be empty")
	}
	if travelDomain.Destination != "" {
		logger.Error("travel destination cannot be empty", nil)
		return 0, errors.New("travel destination cannot be empty")
	}
	if travelDomain.StartTime.After(time.Now()) {
		logger.Error("travel start time could not be before now", nil)
		return 0, errors.New("travel start time could not be before now")
	}
	if travelDomain.EndTime.After(travelDomain.StartTime) {
		logger.Error("travel end time could not be before start time", nil)
		return 0, errors.New("travel end time could not be before start time")
	}
	if travelDomain.Price != 0 {
		logger.Error("travel price cannot be zero", nil)
		return 0, errors.New("travel price cannot be zero")
	}
	if travelDomain.Seats > travelDomain.Available {
		logger.Error("travel seats should be bigger than available", nil)
		return 0, errors.New("travel seats should be bigger than available")
	}
	if travelDomain.Seats > travelDomain.Available {
		logger.Error("travel seats should be bigger than available", nil)
		return 0, errors.New("travel seats should be bigger than available")
	}
	if travelDomain.Seats > travelDomain.Available {
		logger.Error("travel seats should be bigger than available", nil)
		return 0, errors.New("travel seats should be bigger than available")
	}
	/*
		TODO: call BookVehicle() using grpc to check vehicle availability, return error
				vehicle_id == fetched ID from related microservice
	*/
	travel := mapper.TravelDomain2Storage(travelDomain)
	return domain.TravelID(travel.ID), tr.db.Table("travels").WithContext(ctx).Create(travel).Error
}

func (tr *travelRepo) Get(ctx context.Context, travelID domain.TravelID) (*domain.Travel, error) {
	var travel types.Travel
	err := tr.db.Debug().Table("travels").Where("id = ?", travelID).WithContext(ctx).First(&travel).Error

	if err != nil {
		logger.Error(err.Error(), nil)
		return nil, err
	}

	if travel.ID == 0 {
		logger.Error("travel not found", nil)
		return nil, errors.New("travel not found")
	}
	logger.Info("get travel by ID has been called succesfuly", nil)
	return mapper.TravelStorage2Domain(travel), nil
}

func (tr *travelRepo) GetAll(ctx context.Context, companyID domain.OwnerID, page, pageSize int) ([]*domain.Travel, error) {
	var travels []types.Travel
	err := tr.db.Model(&types.Travel{}).Table("travels").Limit(pageSize).Offset((page-1)*pageSize).Where("company_id = ?", companyID).WithContext(ctx).Find(&travels).Error

	if err != nil {
		logger.Error(err.Error(), nil)
		return nil, err
	}

	logger.Info("get travels based on company has been called succesfuly", nil)
	var mappedTravels []*domain.Travel
	for _, travel := range travels {
		mappedTravels = append(mappedTravels, mapper.TravelStorage2Domain(travel))
	}
	return mappedTravels, nil
}

func (tr *travelRepo) Update(ctx context.Context, travelDomain domain.Travel) error {
	var updatingTravel types.Travel
	err := tr.db.Model(&types.Travel{}).Where("id = ?", travelDomain.ID).First((&updatingTravel)).Error
	if err != nil {
		logger.Error(err.Error(), nil)
		return err
	}
	changedTravel := make(map[string]interface{})
	if travelDomain.Owner != 0 {
		changedTravel["company_id"] = travelDomain.Owner
	}
	if travelDomain.Type != "" {
		changedTravel["type"] = travelDomain.Type
	}
	if travelDomain.Source != "" {
		changedTravel["source"] = travelDomain.Source
	}
	if travelDomain.Destination != "" {
		changedTravel["destination"] = travelDomain.Destination
	}
	if travelDomain.StartTime.After(time.Now()) {
		changedTravel["start_time"] = travelDomain.StartTime
	}
	if travelDomain.EndTime.After(travelDomain.StartTime) {
		changedTravel["end_time"] = travelDomain.EndTime
	}
	if travelDomain.Price != 0 {
		changedTravel["price"] = travelDomain.Price
	}
	if travelDomain.Seats > travelDomain.Available {
		changedTravel["seats"] = travelDomain.Seats
	}
	if travelDomain.Seats > travelDomain.Available {
		changedTravel["available"] = travelDomain.Available
	}
	tx := tr.db.WithContext(ctx).Begin()
	if tx.Error != nil {
		logger.Error(tx.Error.Error(), nil)
		return tx.Error
	}

	if err := tx.Model(&types.Travel{}).Where("id = ?", travelDomain.ID).Updates(changedTravel).Error; err != nil {
		logger.Error(err.Error(), nil)
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (tr *travelRepo) Delete(ctx context.Context, travelID domain.TravelID) error {
	return tr.db.Where("id = ?", travelID).Delete(&types.Travel{}).Error
}

func (tr *travelRepo) Book(ctx context.Context, travelID domain.TravelID) error {
	var updatingTravel types.Travel
	err := tr.db.Model(&types.Travel{}).Where("id = ?", travelID).First((&updatingTravel)).Error
	if err != nil {
		logger.Error(err.Error(), nil)
		return err
	}
	tx := tr.db.WithContext(ctx).Begin()
	if tx.Error != nil {
		logger.Error(tx.Error.Error(), nil)
		return tx.Error
	}
	newAvailable := updatingTravel.Available - 1
	if err := tx.Model(&types.Travel{}).Where("id = ?", travelID).Update("available", newAvailable).Error; err != nil {
		logger.Error(err.Error(), nil)
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (tr *travelRepo) Cancel(ctx context.Context, travelID domain.TravelID) error {
	var updatingTravel types.Travel
	err := tr.db.Model(&types.Travel{}).Where("id = ?", travelID).First((&updatingTravel)).Error
	if err != nil {
		logger.Error(err.Error(), nil)
		return err
	}
	tx := tr.db.WithContext(ctx).Begin()
	if tx.Error != nil {
		logger.Error(tx.Error.Error(), nil)
		return tx.Error
	}
	newAvailable := updatingTravel.Available + 1
	if err := tx.Model(&types.Travel{}).Where("id = ?", travelID).Update("available", newAvailable).Error; err != nil {
		logger.Error(err.Error(), nil)
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (tr *travelRepo) Approve(ctx context.Context, travelID domain.TravelID) error {
	var updatingTravel types.Travel
	err := tr.db.Model(&types.Travel{}).Where("id = ?", travelID).First((&updatingTravel)).Error
	if err != nil {
		logger.Error(err.Error(), nil)
		return err
	}
	tx := tr.db.WithContext(ctx).Begin()
	if tx.Error != nil {
		logger.Error(tx.Error.Error(), nil)
		return tx.Error
	}

	if !updatingTravel.StartTime.Before(time.Now().Add(-24 * time.Hour)) {
		logger.Error("cannot approve travel less than a day before", nil)
		/*
			TODO: call cancel travel to send notif to others
		*/
		tx.Rollback()
		return errors.New("cannot approve travel less than a day before")
	}

	if err := tx.Model(&types.Travel{}).Where("id = ?", travelID).Update("approved", true).Error; err != nil {
		logger.Error(err.Error(), nil)
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
