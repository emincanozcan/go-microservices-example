package handlers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/emincanozcan/go-microservices-example/order-service/helpers"
	"github.com/emincanozcan/go-microservices-example/order-service/models"
	"github.com/form3tech-oss/jwt-go"
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

func GetByCurrentUser(c *fiber.Ctx) error {
	var orders []models.Order
	u := c.Locals("user").(*jwt.Token)
	claims := u.Claims.(jwt.MapClaims)
	userID := uint(claims["id"].(float64))
	helpers.DB.Where("user_id = ?", userID).Preload("Items").Find(&orders)
	return c.JSON(fiber.Map{
		"data": orders,
	})
}

type Product struct {
	ID    uint `json:"id" validate:"required"`
	Count uint `json:"count" validate:"required"`
}

func CreateOrder(c *fiber.Ctx) error {
	u := c.Locals("user").(*jwt.Token)
	claims := u.Claims.(jwt.MapClaims)
	userID := uint(claims["id"].(float64))

	var p []Product
	c.BodyParser(&p)

	var totalPrice float32

	// validate data...
	for _, pi := range p {
		v := validator.New()
		if err := v.Struct(&pi); err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "Invalid body"})
		}
	}

	// fetch product data from product service
	var items []models.Item

	for _, ci := range p {
		res, err := http.Get(helpers.Getenv("PRODUCT_SERVICE_BASE_URL") + "api/products/v1/" + strconv.Itoa(int(ci.ID)))
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": "Product service is not accesible"})
		}
		body, _ := ioutil.ReadAll(res.Body)
		var result map[string]interface{}
		json.Unmarshal([]byte(body), &result)

		data := result["data"].(map[string]interface{})

		stock := uint(data["stock"].(float64))
		if ci.Count > stock {
			return c.Status(http.StatusNotAcceptable).JSON(fiber.Map{"message": "Stock problem!"})
		}
		title := data["title"].(string)
		price := float32(data["price"].(float64))
		totalPrice += price * float32(ci.Count)
		items = append(items, models.Item{
			ProductID: ci.ID,
			Title:     title,
			Price:     price,
			Count:     ci.Count,
		})
	}

	// save order
	order := models.Order{
		UserID: userID,
		Price:  totalPrice,
		Items:  items,
	}

	if err := helpers.DB.Create(&order).Error; err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": "Error: " + err.Error()})
	}

	go decreaseStock(items)

	return c.Status(http.StatusCreated).JSON(fiber.Map{"message": "Order is created!", "data": order})
}
func decreaseStock(items []models.Item) {

	client := &http.Client{}
	for _, i := range items {
		json, _ := json.Marshal(map[string]int{"count": int(i.Count)})
		req, _ := http.NewRequest(http.MethodPut,
			helpers.Getenv("PRODUCT_SERVICE_INTERNAL_BASE_URL")+"api/products/v1/"+strconv.Itoa(int(i.ProductID))+"/decrease-stock",
			bytes.NewBuffer(json))
		client.Do(req)
	}
}
