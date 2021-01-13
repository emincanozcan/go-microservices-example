package helpers

import (
	"strings"

	"github.com/emincanozcan/go-microservices-example/product-service/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"syreclabs.com/go/faker"
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
	// seed()
}

func migrate() {
	DB.AutoMigrate(&models.Product{})
}

func seed() {
	for i := 0; i < 1000; i++ {
		product := models.Product{
			Title:       strings.Title(faker.Lorem().Sentence(3)),
			Stock:       uint(faker.Number().NumberInt(3)),
			Price:       faker.Commerce().Price(),
			Description: strings.Join(faker.Lorem().Sentences(8), "."),
		}
		DB.Create(&product)
	}
}
