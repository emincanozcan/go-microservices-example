package handlers

import (
	"github.com/emincanozcan/go-microservices-example/order-service/helpers"
	"github.com/emincanozcan/go-microservices-example/order-service/models"
	"github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
)

func GetByCurrentUser(c *fiber.Ctx) error {
	var orders []models.Order
	u := c.Locals("user").(*jwt.Token)
	claims := u.Claims.(jwt.MapClaims)
	userID := uint(claims["id"].(float64))
	helpers.DB.Preload("Items").Where("user_id = ?", userID).Find(&orders)
	return c.JSON(fiber.Map{
		"data": orders,
	})
}

type Product struct {
	ID    uint `json:"id" validate:"required"`
	Count uint `json:"count" validate:"required"`
}

func CreateOrder(c *fiber.Ctx) error {
	// u := c.Locals("user").(*jwt.Token)
	// claims := u.Claims.(jwt.MapClaims)
	// userID := uint(claims["id"].(float64))

	// var p []Product
	// c.BodyParser(&p)

	// var pIds []int
	// for _, pi := range p {
	// 	v := validator.New()
	// 	pIds = append(pIds, int(pi.ID))
	// 	err := v.Struct(&pi)
	// 	if err != nil {
	// 		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "Invalid body"})
	// 	}
	// }

	// var Items []models.Item

	// for _, pi := range p {
	// 	pId := int(pi.ID)
	// 	r, err := http.Get("http://product-service:3000/products/" + strconv.Itoa(pId))
	// 	if err != nil {
	// 		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "Product service is not accesible"})
	// 	}
	// 	body, _ := ioutil.ReadAll(r.Body)
	// 	var result map[string]interface{}
	// 	json.Unmarshal([]byte(body), &result)

	// 	data := result["data"].(map[string]string)
	// 	id, _ := strconv.Atoi(data["id"])
	// 	title := data["title"]
	// 	price, _ := strconv.ParseFloat(data["price"], 32)

	// 	Items = append(Items, models.Item{
	// 		ID:    uint(id),
	// 		Title: title,
	// 		Price: float32(price),
	// 		Count: pi.Count,
	// 	})
	// }
	return nil
}
