package routes_v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/matinkaboli/goods-golang/routes/v1/users"
)

func Routes(app *fiber.App) {
	v1 := app.Group("/v1")
	usersGroup := v1.Group("/users")

	usersGroup.Post("/", users.RegisterUsersHandler)
	usersGroup.Post("/login", users.LoginUsersHandler)
}
