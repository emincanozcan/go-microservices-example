package helpers

import (
	"github.com/emincanozcan/go-microservices-example/order-service/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DbConnect() {
	dsn := "root:password@tcp(order-database)/orders?charset=utf8mb4"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Database connection error!")
	}
	DB = db
	migrate()
}
func migrate() {
	DB.AutoMigrate(&models.Order{})
	DB.AutoMigrate(&models.Item{})
}
