package helpers

import (
	"github.com/bxcodec/faker/v3"
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
	// seed()
}

func migrate() {
	DB.AutoMigrate(&models.User{})
}

func seed() {
	user := models.User{
		Name:     "Admin Account",
		Email:    "admin@admin.com",
		Password: "adminpass",
		Type:     9,
	}
	DB.Create(&user)
	for i := 0; i < 100; i++ {
		user := models.User{
			Name:     faker.Name(),
			Email:    faker.Email(),
			Password: faker.Password(),
		}
		DB.Create(&user)
	}
}
