package main

import (
	"net/http"
	"sync"

	"github.com/emincanozcan/go-microservices-example/product-service/handlers"
	"github.com/emincanozcan/go-microservices-example/product-service/helpers"
	"github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
)

var wg = &sync.WaitGroup{}

func main() {
	wg.Add(2)
	helpers.DatabaseConnect()
	go initGlobalService()
	go initInternalService()
	wg.Wait()
}
func initGlobalService() {
	defer wg.Done()
	app := fiber.New()
	app.Get("/api/products/v1", handlers.GetProducts)
	app.Get("/api/products/v1/:id", handlers.GetProduct)

	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte(helpers.Getenv("JWT_KEY")),
	}))

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
	app.Post("/api/products/v1", handlers.CreateProduct)
	app.Put("/api/products/v1/:id", handlers.UpdateProduct)
	app.Delete("/api/products/v1/:id", handlers.DeleteProduct)
	app.Listen(":3000")
}
func initInternalService() {
	defer wg.Done()
	app := fiber.New()
	app.Put("/api/products/v1/:id/decrease-stock", handlers.DecreaseStockOfProduct)
	app.Listen(":3001")
}
