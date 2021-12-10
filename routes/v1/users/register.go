package users

import (
	"crypto/sha256"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/matinkaboli/goods-golang/database"
	"github.com/matinkaboli/goods-golang/models"
	"github.com/matinkaboli/goods-golang/utils"
)

type RegisterBody struct {
	Name     string `validate:"required,min=5"`
	Phone    string `validate:"required,min=10,max=13"`
	Password string `validate:"required,min=8"`
}

func RegisterUsersHandler(c *fiber.Ctx) error {
	body := new(RegisterBody)

	if err := c.BodyParser(body); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Invalid Body")
	}

	errors := utils.ValidateStruct(*body)

	if errors != nil {
		return utils.NewError(c, fiber.StatusBadRequest, errors)
	}

	hashed := sha256.Sum256([]byte(body.Password))
	hashed2 := fmt.Sprintf("%x", hashed)

	db := database.DB

	user := &models.User{
		Name:     body.Name,
		Phone:    body.Phone,
		Password: hashed2,
	}

	if dbc := db.Create(user); dbc.Error != nil {
		if utils.IsUniqueViolation(dbc.Error) {
			return fiber.NewError(fiber.StatusConflict, "Phone is duplicated")
		}
	}

	return c.JSON(&body)
}
