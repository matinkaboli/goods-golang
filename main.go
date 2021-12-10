package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/matinkaboli/goods-golang/database"
	routes_v1 "github.com/matinkaboli/goods-golang/routes/v1"
)

func main() {
	app := fiber.New(fiber.Config{})
	database.Connect()

	routes_v1.Routes(app)

	log.Fatal(app.Listen(":4567"))
}
