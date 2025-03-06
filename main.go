package main

import (
	"log"

	"github.com/akekawich36/go-authen/configs"
	"github.com/gofiber/fiber/v2"
)

func main() {
	configs.DatabaseConnect()

	app := fiber.New()

	log.Fatal(app.Listen(":9000"))
}
