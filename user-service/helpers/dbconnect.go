package helpers

import (
	"github.com/emincanozcan/go-microservices-example/user-service/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DbConnect() {
	dsn := "root:password@tcp(user-database)/users?charset=utf8mb4"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("DATABASE CONNECTION FAULT")
	}
	DB = db
	migrate()
}

func migrate() {
	DB.AutoMigrate(&models.User{})
}
