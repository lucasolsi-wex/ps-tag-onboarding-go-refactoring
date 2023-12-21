package routes

import (
	"github.com/gin-gonic/gin"
	"lucasolsi-wex/ps-tag-onboarding-go/src/controller"
)

func InitRoutes(rg *gin.RouterGroup) {
	rg.GET("/user/:userId", controller.GetUserById)
	rg.POST("/user", controller.CreateUser)
}
