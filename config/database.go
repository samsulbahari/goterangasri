package config

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func SetDatabase() (*gorm.DB, error) {
	var dsn = os.Getenv("DB_CONFIG")
	var Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{NamingStrategy: schema.NamingStrategy{
		SingularTable: true,
	}})

	return Db, err

}
