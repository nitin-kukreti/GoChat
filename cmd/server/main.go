package main

import (
	"github.com/nitin-kukreti/GoChat/internal/config"
	"github.com/nitin-kukreti/GoChat/internal/interface/controller"
	"github.com/nitin-kukreti/GoChat/internal/interface/storage"
	"github.com/nitin-kukreti/GoChat/internal/routes"
	"github.com/nitin-kukreti/GoChat/internal/server"
	"github.com/nitin-kukreti/GoChat/internal/usecase"
)

func main() {
	db := config.ConnectDB()
	defer db.Close()

	// Setup dependencies
	userRepo := storage.NewUserRepository(db)
	groupRepo := storage.NewGroupRepository(db)

	userUC := usecase.NewUserUseCase(userRepo)
	groupUC := usecase.NewGroupUseCase(groupRepo)

	userHandler := controller.NewUserHandler(userUC)
	groupHandler := controller.NewGroupHandler(groupUC)

	app := server.NewServer()

	// Register all routes grouped by domain
	routes.RegisterAllRoutes(app, userHandler, groupHandler)

	app.Listen(":8080")
}
