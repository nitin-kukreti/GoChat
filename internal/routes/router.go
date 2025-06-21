package routes

import (
	"github.com/nitin-kukreti/GoChat/internal/interface/controller"
	"github.com/nitin-kukreti/GoChat/internal/server"
)

func RegisterAllRoutes(app server.App, userHandler *controller.UserHandler, groupHandler *controller.GroupHandler) {
	// grp:=app
	RegisterUserRoutes(app, userHandler)
	RegisterGroupRoutes(app, groupHandler)
}
