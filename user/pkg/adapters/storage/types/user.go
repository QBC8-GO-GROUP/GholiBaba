package types

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName string
	LastName  string
	Phone     string
	Password  string
	WalletID  uuid.UUID
	Role      Role
}

type Role string

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
