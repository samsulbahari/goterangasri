package config

import (
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func SetDatabase() (*gorm.DB, error) {
	var dsn = os.Getenv("DB_CONFIG")
	var Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{NamingStrategy: schema.NamingStrategy{
		SingularTable: true,
	}})

	return Db, err

}
