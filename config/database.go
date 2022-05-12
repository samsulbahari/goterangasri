package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func SetDatabase() (*gorm.DB, error) {
	var dsn = "host=ec2-54-164-40-66.compute-1.amazonaws.com user=tgegltttehbhck password=536dd6f2a4958f994bd37c27c64d17dc6026a13452acde53f59bc901dc81d825 dbname=dbuvp1i6mcmshg port=5432 sslmode=require TimeZone=Asia/Shanghai"
	var Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{NamingStrategy: schema.NamingStrategy{
		SingularTable: true,
	}})

	return Db, err

}
