package postgres

import (
	"fmt"
	"log"

	"github.com/QBC8-GO-GROUP/GholiBaba/pkg/adapters/storage/types"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DBConnOptions struct {
	Company string
	Pass    string
	Host    string
	Port    uint
	DBName  string
	Schema  string
}

func (o DBConnOptions) PostgresDSN() string {
	return fmt.Sprintf("host=%s port=%d company=%s password=%s dbname=%s search_path=%s sslmode=disable",
		o.Host, o.Port, o.Company, o.Pass, o.DBName, o.Schema)
}

func NewPsqlGormConnection(opt DBConnOptions) (*gorm.DB, error) {
	return gorm.Open(postgres.Open(opt.PostgresDSN()), &gorm.Config{
		Logger: logger.Default,
	})
}
func GormMigrations(db *gorm.DB) {

	err := db.AutoMigrate(
		&types.Company{},
	)
	if err != nil {
		log.Fatalf("failed to migrate models: %v", err)
	}
}
