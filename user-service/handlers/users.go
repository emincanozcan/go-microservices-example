package handlers

import (
	"net/http"
	"time"

	"github.com/emincanozcan/go-microservices-example/user-service/helpers"
	"github.com/emincanozcan/go-microservices-example/user-service/models"
	jwt "github.com/form3tech-oss/jwt-go"
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

type UserRegisterStruct struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func Register(c *fiber.Ctx) error {
	var u UserRegisterStruct
	c.BodyParser(&u)

	v := validator.New()
	err := v.Struct(&u)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(map[string]string{
			"message": "Data is not valid",
		})

	}

	var User models.User
	User.Email = u.Email
	User.Name = u.Name
	User.Password = u.Password

	helpers.DB.Create(&User)

	return c.Status(201).JSON(map[string]interface{}{
		"data": User,
	})
}

type UserLoginStruct struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func Login(c *fiber.Ctx) error {
	var u UserLoginStruct
	c.BodyParser(&u)

	v := validator.New()
	err := v.Struct(&u)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(map[string]string{
			"message": "Bad request",
		})

	}

	var user models.User
	res := helpers.DB.Where("email", u.Email).First(&user)
	if res.RowsAffected < 1 {
		return c.Status(fiber.StatusNotFound).JSON(map[string]string{
			"message": "Email not found",
		})
	}

	if user.Password != u.Password {
		return c.Status(fiber.StatusNotFound).JSON(map[string]string{
			"message": "Wrong password",
		})

	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = user.ID
	claims["isAdmin"] = user.Type == 9
	claims["time"] = time.Now()

	t, _ := token.SignedString([]byte(helpers.Getenv("JWT_KEY")))

	return c.JSON(map[string]map[string]string{
		"data": {
			"token": t,
		},
	})
}

func CurrentUser(c *fiber.Ctx) error {
	u := c.Locals("user").(*jwt.Token)
	claims := u.Claims.(jwt.MapClaims)
	id := claims["id"]

	var user models.User
	res := helpers.DB.Where("id", id).First(&user)
	if res.RowsAffected < 1 {
		return c.Status(http.StatusUnauthorized).JSON(map[string]string{
			"message": "User not found",
		})
	}
	return c.JSON(fiber.Map{"data": user})
}
