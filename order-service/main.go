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
	app.Get("/orders", handlers.GetByCurrentUser)
	app.Post("/orders", handlers.CreateOrder)
	app.Listen(":3000")
}
