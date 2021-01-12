package helpers

import (
	"github.com/emincanozcan/go-microservices-example/product-service/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DatabaseConnect() {
	dsn := "root:password@tcp(product-database)/products?charset=utf8mb4"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("DATABASE CONNECTION ERROR!")
	}
	DB = db
	migrate()
}

func migrate() {
	DB.AutoMigrate(&models.Product{})
}
