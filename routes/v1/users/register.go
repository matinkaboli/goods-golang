package users

import (
	"crypto/sha256"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/matinkaboli/goods-golang/database"
	"github.com/matinkaboli/goods-golang/models"
	"github.com/matinkaboli/goods-golang/utils"
)

func RegisterUsersHandler(c *fiber.Ctx) error {
	name := c.FormValue("name")
	phone := c.FormValue("phone")
	password := c.FormValue("password")

	if name == "" || phone == "" || password == "" {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	hashed := sha256.Sum256([]byte(password))
	hashed2 := fmt.Sprintf("%x", hashed)

	db := database.DB

	user := &models.User{
		Name:     name,
		Phone:    phone,
		Password: hashed2,
	}

	if dbc := db.Create(user); dbc.Error != nil {
		if utils.IsUniqueViolation(dbc.Error) {
			return c.SendString("Duplicated")
		}
	}

	return c.JSON(&user)
}
