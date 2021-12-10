package users

import (
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

	// Parse JSON from Body Parser and fill body
	if err := c.BodyParser(body); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Invalid Body")
	}

	// Validate body from RegisterBody
	if errors := utils.ValidateStruct(*body); errors != nil {
		return utils.NewError(c, fiber.StatusBadRequest, errors)
	}

	db := database.DB

	user := &models.User{
		Name:     body.Name,
		Phone:    body.Phone,
		Password: utils.HashString(body.Password),
	}

	if dbc := db.Create(user); dbc.Error != nil {
		if utils.IsUniqueViolation(dbc.Error) {
			return fiber.NewError(fiber.StatusConflict, "Phone is duplicated")
		}
	}

	return c.JSON(&body)
}
