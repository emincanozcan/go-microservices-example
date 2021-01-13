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
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
}

func CreateProduct(c *fiber.Ctx) error {
	var p createProductStructure
	c.BodyParser(&p)
	v := validator.New()
	err := v.Struct(p)
	if err != nil {
		return c.JSON(map[string]string{
			"message": "Invalid data",
		})

	}

	product := models.Product{Title: p.Title, Description: p.Description}
	helpers.DB.Create(&product)
	return c.Status(201).JSON(product)
}

type updateProductStructure struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func UpdateProduct(c *fiber.Ctx) error {
	var p createProductStructure
	c.BodyParser(&p)
	v := validator.New()
	err := v.Struct(p)
	if err != nil {
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
