package users

import (
	"github.com/gofiber/fiber/v2"
	"github.com/matinkaboli/goods-golang/database"
	"github.com/matinkaboli/goods-golang/models"
)

func RegisterUsersHandler(c *fiber.Ctx) error {
	name := c.FormValue("name")
	phone := c.FormValue("phone")
	password := c.FormValue("password")

	if name == "" || phone == "" || password == "" {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	db := database.DB

	db.Create(&models.User{
		Name:     name,
		Phone:    phone,
		Password: password,
	})

	return c.JSON(fiber.Map{
		"ok": true,
	})
}
