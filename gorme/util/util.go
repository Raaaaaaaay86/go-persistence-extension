package util

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func CreateGormPostgreSqlConnection(config *gorm.Config) (*gorm.DB, error) {
	driver := postgres.New(postgres.Config{
		DSN: "host=localhost port=5439 user=gpe_test password=changeme dbname=gpe sslmode=disable",
	})

	return gorm.Open(driver, config)
}
