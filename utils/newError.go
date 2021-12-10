package utils

import "github.com/gofiber/fiber/v2"

func NewError(c *fiber.Ctx, code int, errors []*ErrorResponse) error {
	return c.Status(code).JSON(errors)
}
