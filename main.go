package main

import (
	"log"

	"github.com/akekawich36/go-authen/internal/delivery/http/route"
	"github.com/akekawich36/go-authen/internal/domain/models"
	"github.com/akekawich36/go-authen/internal/infrastructure/database"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
)

func main() {
	db, err := database.DatabaseConnect()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	db.AutoMigrate(&models.User{})

	app := fiber.New()

	route.SetupRoute(app, db)

	app.Use(helmet.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	log.Fatal(app.Listen(":9000"))
}
