package routes;


import (
	"github.com/nitin-kukreti/GoChat/internal/interface/controller"
	"github.com/nitin-kukreti/GoChat/internal/server"
)

func RegisterUserRoutes(app server.App, userHandler *controller.UserHandler) {
	userRouter := app.Group("/api/users")
	userRouter.POST("/", userHandler.CreateUserHandler);
	userRouter.GET("/{id}", userHandler.GetUserById);
}
