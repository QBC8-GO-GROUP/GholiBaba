package postgres

import (
	"fmt"

	"github.com/QBC8-GO-GROUP/GholiBaba/pkg/adapters/storage/types"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DBConnOptions struct {
	User   string
	Pass   string
	Host   string
	Port   uint
	DBName string
	Schema string
}

func (o DBConnOptions) PostgresDSN() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s search_path=%s sslmode=disable",
		o.Host, o.Port, o.User, o.Pass, o.DBName, o.Schema)
}

//	func NewPsqlGormConnection(opt DBConnOptions) (*gorm.DB, error) {
//		return gorm.Open(postgres.Open(opt.PostgresDSN()), &gorm.Config{
//			Logger: logger.Discard,
//		})
//	}
func NewPsqlGormConnection(opt DBConnOptions) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(opt.PostgresDSN()), &gorm.Config{Logger: logger.Discard})
	if err != nil {
		fmt.Printf("Database connection failed with DSN: %v\n", db)
		return nil, err
	}

	if err := migrate(db); err != nil {
		fmt.Printf("migrations failed: %v\n", err.Error())
		return nil, err
	}
	return db, nil
}

func migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&types.User{},
	)
}
