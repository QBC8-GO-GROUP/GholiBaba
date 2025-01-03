package mapper

import (
	"github.com/QBC8-GO-GROUP/GholiBaba/internal/user/domain"
	"github.com/QBC8-GO-GROUP/GholiBaba/pkg/adapters/storage/types"

	"gorm.io/gorm"
)

func UserDomain2Storage(userDomain domain.User) *types.User {
	return &types.User{
		Model: gorm.Model{
			ID:        uint(userDomain.ID),
			CreatedAt: userDomain.CreatedAt,
			DeletedAt: gorm.DeletedAt(ToNullTime(userDomain.DeletedAt)),
		},
		FirstName: userDomain.FirstName,
		LastName:  userDomain.LastName,
		Phone:     string(userDomain.Phone),
		Password:  userDomain.Password,
		// WalletID:  userDomain.WalletID,
		Role: types.Role(userDomain.Role),
	}
}

func UserStorage2Domain(user types.User) *domain.User {
	return &domain.User{
		ID:        domain.UserID(user.ID),
		CreatedAt: user.CreatedAt,
		DeletedAt: user.DeletedAt.Time,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Phone:     domain.Phone(user.Phone),
		Password:  user.Password,
		// WalletID:  user.WalletID,
		Role: domain.Role(user.Role),
	}
}
