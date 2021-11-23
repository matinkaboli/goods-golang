package users

import "github.com/gofiber/fiber/v2"

func LoginUsersHandler(c *fiber.Ctx) error {
	name := c.FormValue("name")
	phone := c.FormValue("phone")
	password := c.FormValue("password")

	if name == "" || phone == "" || password == "" {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	return c.SendString("I am login route")
}
