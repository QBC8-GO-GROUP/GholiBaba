package storage

import (
	"context"
	"errors"

	"github.com/QBC8-GO-GROUP/GholiBaba/internal/user/domain"
	"github.com/QBC8-GO-GROUP/GholiBaba/internal/user/port"
	"github.com/QBC8-GO-GROUP/GholiBaba/pkg/adapters/storage/mapper"
	"github.com/QBC8-GO-GROUP/GholiBaba/pkg/adapters/storage/types"
	"gorm.io/gorm"
)

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) port.Repo {
	return &userRepo{db}
}

func (r *userRepo) Create(ctx context.Context, userDomain domain.User) (domain.UserID, error) {
	user := mapper.UserDomain2Storage(userDomain)
	return domain.UserID(user.ID), r.db.Table("users").WithContext(ctx).Create(user).Error
}

func (r *userRepo) GetByFilter(ctx context.Context, filter *domain.UserFilter) (*domain.User, error) {
	var user types.User

	q := r.db.Table("users").Debug().WithContext(ctx)

	if filter.ID > 0 {
		q = q.Where("id = ?", filter.ID)
	}

	if len(filter.Phone) > 0 {
		q = q.Where("phone = ?", filter.Phone)
	}

	err := q.First(&user).Error

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	if user.ID == 0 {
		return nil, nil
	}

	return mapper.UserStorage2Domain(user), nil
}

func (r *userRepo) Update(ctx context.Context, user domain.User) error {

	var existingUser domain.User

	if err := r.db.Where("id = ?", user.ID).First(&existingUser).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("user not found")
		}
		return err
	}

	if err := r.db.Model(&existingUser).Updates(user).Error; err != nil {
		return err
	}

	return nil
}

func (r *userRepo) GetById(ctx context.Context, id domain.UserID) (domain.User, error) {

	var foundUser domain.User

	if err := r.db.Where("id = ?", id).First(&foundUser).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return foundUser, errors.New("user not found")
		}
		return foundUser, err
	}

	return foundUser, nil
}
