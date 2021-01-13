package helpers

import (
	"strings"

	"github.com/emincanozcan/go-microservices-example/order-service/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"syreclabs.com/go/faker"
)

var DB *gorm.DB

func DbConnect() {
	dsn := "root:password@tcp(order-database)/orders?charset=utf8mb4&parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Database connection error!")
	}
	DB = db
	migrate()
	// seed()
}
func migrate() {
	DB.AutoMigrate(&models.Order{})
	DB.AutoMigrate(&models.Item{})
}

func seed() {
	for i := 0; i < 5000; i++ {
		var items []models.Item
		var totalPrice float32
		for j := 0; j < 3; j++ {
			p := faker.Commerce().Price()
			c := float32(faker.Number().NumberInt(1))
			totalPrice += p * c
			items = append(items, models.Item{
				Price:     p,
				ProductID: uint(faker.Number().NumberInt(3)),
				Count:     uint(c),
				Title:     strings.Title(faker.Lorem().Sentence(3)),
			})
		}
		order := models.Order{
			Items:  items,
			Price:  totalPrice,
			UserID: uint(faker.Number().NumberInt(2)),
		}
		DB.Create(&order)
	}
}
