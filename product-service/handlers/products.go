package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/emincanozcan/go-microservices-example/product-service/helpers"
	"github.com/emincanozcan/go-microservices-example/product-service/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"github.com/go-playground/validator/v10"
)

func GetProducts(c *fiber.Ctx) error {
	var p []models.Product
	helpers.DB.Scopes(helpers.Paginate(c)).Find(&p)
	return c.Status(http.StatusOK).JSON(map[string]interface{}{
		"data": p,
	})
}

type createProductStructure struct {
	Title       string  `json:"title" validate:"required"`
	Description string  `json:"description" validate:"required"`
	Stock       uint    `json:"stock" validate:"required"`
	Price       float32 `json:"price" validate:"required"`
}

func isProductValidProductInput(c *fiber.Ctx) (createProductStructure, bool) {
	var p createProductStructure
	c.BodyParser(&p)
	v := validator.New()
	err := v.Struct(p)
	if err != nil {
		return createProductStructure{}, false
	}
	return p, true
}

func CreateProduct(c *fiber.Ctx) error {
	p, ok := isProductValidProductInput(c)
	if !ok {
		return c.JSON(map[string]string{
			"message": "Invalid data",
		})
	}
	product := models.Product{Title: p.Title, Description: p.Description, Stock: p.Stock, Price: p.Price}
	helpers.DB.Create(&product)
	return c.Status(201).JSON(product)
}

func UpdateProduct(c *fiber.Ctx) error {
	p, ok := isProductValidProductInput(c)
	if !ok {
		return c.JSON(map[string]string{
			"message": "Invalid data",
		})
	}

	product := models.Product{}
	var pID = c.Params("id")
	r := helpers.DB.Where("id", pID).First(&product)
	if r.RowsAffected < 1 {
		return c.JSON(map[string]string{
			"message": "Product is not found",
		})
	}

	product.Title = p.Title
	product.Description = p.Description
	product.Price = p.Price
	product.Stock = p.Stock
	helpers.DB.Updates(&product)
	return c.JSON(map[string]interface{}{
		"data": product,
	})
}

func DeleteProduct(c *fiber.Ctx) error {
	pID, _ := strconv.Atoi(c.Params("id"))
	p := models.Product{ID: uint(pID)}
	r := helpers.DB.Delete(&p)
	if r.RowsAffected < 1 {
		return c.JSON(map[string]string{
			"message": "Product is not found",
		})

	}
	return c.JSON(map[string]interface{}{
		"message": "Deleted",
	})
}

func GetProduct(c *fiber.Ctx) error {
	pID, _ := strconv.Atoi(c.Params("id"))
	p := models.Product{}
	r := helpers.DB.Where("id", pID).First(&p)
	if r.RowsAffected < 1 {
		return c.JSON(map[string]string{
			"message": "Product is not found",
		})

	}
	return c.JSON(map[string]interface{}{
		"data": p,
	})
}

func DecreaseStockOfProduct(c *fiber.Ctx) error {
	var body map[string]interface{}
	json.Unmarshal(c.Body(), &body)
	count := int(body["count"].(float64))

	pID, _ := strconv.Atoi(c.Params("id"))
	helpers.DB.Model(&models.Product{}).Where("id", pID).UpdateColumn("stock", gorm.Expr("stock - ?", count))
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "OK"})
}
