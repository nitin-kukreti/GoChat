package routes

import (
	"github.com/nitin-kukreti/GoChat/internal/interface/controller"
	"github.com/nitin-kukreti/GoChat/internal/server"
)

func RegisterGroupRoutes(app server.App, groupHandler *controller.GroupHandler) {
	groupRouter:=app.Group("api/v1/group")
	groupRouter.POST("/", groupHandler.CreateGroupHandler)
	groupRouter.POST("/add-user", groupHandler.AddUserToGroup)
}
