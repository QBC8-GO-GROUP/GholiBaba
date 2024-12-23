package types

import (
	"gorm.io/gorm"
)

type Company struct {
	gorm.Model
	FirstName string
	LastName  string
	Phone     string
	Password  string
}
