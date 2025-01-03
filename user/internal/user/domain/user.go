package domain

import (
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/QBC8-GO-GROUP/GholiBaba/api/pb"

	"github.com/QBC8-GO-GROUP/GholiBaba/pkg/conv"
	"github.com/google/uuid"
)

// value object
type (
	UserID uint
	Phone  string
	Role   string
)

func (p Phone) IsValid() bool {
	// todo regex
	return true
}

type User struct {
	ID        UserID
	CreatedAt time.Time
	DeletedAt time.Time
	FirstName string
	LastName  string
	Password  string
	Phone     Phone
	Role      Role
	WalletID  uuid.UUID
}

const (
	Admin                   Role = "admin"
	RegularUser             Role = "regular_user"
	BusTechnicalTeam        Role = "bus_technical_team"
	CruiseShipTechnicalTeam Role = "cruise_ship_technical_team"
	AirplaneTechnicalTeam   Role = "airplane_technical_team"
	TrainTechnicalTeam      Role = "train_technical_team"
	TransportationCompanies Role = "transportation_companies"
	TravelAgencies          Role = "travel_agencies"
	Hotels                  Role = "hotels"
	RealOwnerOfVehicles     Role = "real_owner_of_vehicles"
)

func (u *User) Validate() error {
	if !u.Phone.IsValid() {
		return errors.New("phone is not valid")
	}
	return nil
}

func (u *User) PasswordIsCorrect(pass string) bool {
	// return NewPassword(pass) == u.Password
	return pass == u.Password
}

func NewPassword(pass string) string {
	h := sha256.New()
	h.Write(conv.ToBytes(pass))
	return base64.URLEncoding.EncodeToString(h.Sum(nil))
}

type UserFilter struct {
	ID    UserID
	Phone string
}

func (f *UserFilter) IsValid() bool {
	f.Phone = strings.TrimSpace(f.Phone)
	return f.ID > 0 || len(f.Phone) > 0
}

func AllowedRoles() []Role {
	return []Role{
		RegularUser,
		BusTechnicalTeam,
		CruiseShipTechnicalTeam,
		AirplaneTechnicalTeam,
		TrainTechnicalTeam,
		TransportationCompanies,
		TravelAgencies,
		Hotels,
		RealOwnerOfVehicles,
	}
}

// IsValidRole checks if the given role is in the allowed list
func IsValidRole(role Role) bool {
	for _, allowedRole := range AllowedRoles() {
		if role == allowedRole {
			return true
		}
	}
	return false
}
func MapProtoRoleToDomain(protoRole pb.Role) (Role, error) {
	switch protoRole {
	case pb.Role_ROLE_ADMIN:
		return Admin, nil
	case pb.Role_ROLE_REGULAR_USER:
		return RegularUser, nil
	case pb.Role_ROLE_BUS_TECHNICAL_TEAM:
		return BusTechnicalTeam, nil
	case pb.Role_ROLE_CRUISE_SHIP_TECHNICAL_TEAM:
		return CruiseShipTechnicalTeam, nil
	case pb.Role_ROLE_AIRPLANE_TECHNICAL_TEAM:
		return AirplaneTechnicalTeam, nil
	case pb.Role_ROLE_TRAIN_TECHNICAL_TEAM:
		return TrainTechnicalTeam, nil
	case pb.Role_ROLE_TRANSPORTATION_COMPANIES:
		return TransportationCompanies, nil
	case pb.Role_ROLE_TRAVEL_AGENCIES:
		return TravelAgencies, nil
	case pb.Role_ROLE_HOTELS:
		return Hotels, nil
	case pb.Role_ROLE_REAL_OWNER_OF_VEHICLES:
		return RealOwnerOfVehicles, nil
	default:
		return "", fmt.Errorf("invalid role provided")
	}
}
