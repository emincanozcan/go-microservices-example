package main

import (
	"net/http"

	"github.com/emincanozcan/go-microservices-example/product-service/handlers"
	"github.com/emincanozcan/go-microservices-example/product-service/helpers"
	"github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
)

func main() {
	helpers.DatabaseConnect()
	app := fiber.New()
	app.Get("/products", handlers.GetProducts)
	app.Get("/products/:id", handlers.GetProduct)

	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte(helpers.Getenv("JWT_KEY")),
	}))

	adminRoutes(app)
	app.Listen(":3000")
}

func adminRoutes(app *fiber.App) {
	app.Use(func(c *fiber.Ctx) error {
		u := c.Locals("user").(*jwt.Token)
		claims := u.Claims.(jwt.MapClaims)
		if !claims["isAdmin"].(bool) {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		}
		return c.Next()
	})
	app.Post("/products", handlers.CreateProduct)
	app.Put("/products/:id", handlers.UpdateProduct)
	app.Delete("/products/:id", handlers.DeleteProduct)
}
