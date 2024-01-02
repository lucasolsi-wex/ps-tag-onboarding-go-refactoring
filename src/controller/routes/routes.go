package routes

import (
	"github.com/gin-gonic/gin"
	"lucasolsi-wex/ps-tag-onboarding-go/src/controller"
)

func InitRoutes(rg *gin.RouterGroup, controllerInterface controller.UserControllerInterface) {
	rg.GET("/user/:userId", controllerInterface.FindUserById)
	rg.POST("/user", controllerInterface.CreateUser)
}
