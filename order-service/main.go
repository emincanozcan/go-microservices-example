package main

import (
	"github.com/emincanozcan/go-microservices-example/order-service/handlers"
	"github.com/emincanozcan/go-microservices-example/order-service/helpers"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
)

func main() {
	helpers.DbConnect()
	app := fiber.New()
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte(helpers.Getenv("JWT_KEY")),
	}))
	app.Get("/api/orders/v1", handlers.GetByCurrentUser)
	app.Post("/api/orders/v1", handlers.CreateOrder)
	app.Listen(":3000")
}
