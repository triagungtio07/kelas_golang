package database

import (
	"fmt"

	"github.com/triagungtio07/golang_fiber/config/env"
	"github.com/triagungtio07/golang_fiber/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	GormDB *gorm.DB
)

func Load() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true&loc=Local",
		env.DbUser, env.DbPass, env.DbHost, env.DbPort, env.DbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	GormDB = db

	if env.DbAutoMigrate {
		GormDB.AutoMigrate(
			&models.Product{},
		)
	}
}
