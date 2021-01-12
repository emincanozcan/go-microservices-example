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
	app.Post("/users/register", handlers.Register)
	app.Post("/users/login", handlers.Login)
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte(helpers.Getenv("JWT_KEY")),
	}))
	app.Get("/users/current-user", handlers.CurrentUser)
	app.Listen(":3000")
}
