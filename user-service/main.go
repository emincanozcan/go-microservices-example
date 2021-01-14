package main

import (
	"github.com/emincanozcan/go-microservices-example/user-service/handlers"
	"github.com/emincanozcan/go-microservices-example/user-service/helpers"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
)

func main() {
	helpers.DbConnect()
	app := fiber.New()
	app.Post("/api/users/v1/register", handlers.Register)
	app.Post("/api/users/v1/login", handlers.Login)
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte(helpers.Getenv("JWT_KEY")),
	}))
	app.Get("/api/users/v1/current-user", handlers.CurrentUser)
	app.Listen(":3000")
}
