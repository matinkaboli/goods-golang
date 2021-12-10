package utils

import "github.com/gofiber/fiber/v2"

func NewError(c *fiber.Ctx, code int, errors []*ErrorResponse) error {
	e := &fiber.Error{
		Code:    code,
		Message: "HIHIHIHIH",
	}

	return c.Status(e.Code).JSON(errors)
}
