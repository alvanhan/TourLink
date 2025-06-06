package main

import (
	"tourlink/internal/user/delivery/http"
	"tourlink/internal/user/infrastructure"
	"tourlink/internal/user/usecase"
	"tourlink/pkg/database"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	userRepo := infrastructure.NewUserRepository(database.DB)
	userService := usecase.NewUserService(userRepo)
	userHandler := http.NewUserHandler(userService)

	// User routes group
	userGroup := app.Group("/api/users")
	userHandler.RegisterRoutes(userGroup)
}
