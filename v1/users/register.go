package users

import "github.com/gofiber/fiber/v2"

func RegisterUsers(users fiber.Router) {
	users.Post("/", func(c *fiber.Ctx) error {
		return c.SendString("I am in a new file")
	})
}
