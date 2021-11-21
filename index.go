package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/matinkaboli/goods-golang/v1/users"
)

func main() {
	app := fiber.New()

	api := app.Group("/api")
	v1 := api.Group("/v1")
	usersGroup := v1.Group("/users")

	users.RegisterUsers(usersGroup)

	log.Fatal(app.Listen(":4567"))
}
