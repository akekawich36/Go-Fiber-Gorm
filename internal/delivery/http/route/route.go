package route

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"github.com/akekawich36/go-authen/internal/delivery/http/handler"
	"github.com/akekawich36/go-authen/internal/domain/usecase"
)

func SetupRoute(app *fiber.App, db *gorm.DB) {
	api := app.Group("/api")
	auth := api.Group("/auth")

	authUseCase := usecase.NewAuthUseCase(userRepo)
	authHandler := handler.NewAuthHandler(authUseCase)
	auth.Post("/register", authHandler.Register)
}
