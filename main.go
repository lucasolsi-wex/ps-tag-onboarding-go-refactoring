package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
	"lucasolsi-wex/ps-tag-onboarding-go/src/controller"
	"lucasolsi-wex/ps-tag-onboarding-go/src/controller/routes"
	"lucasolsi-wex/ps-tag-onboarding-go/src/service"
)

func main() {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		return
	}

	userService := service.NewUserDomainService()
	userController := controller.NewUserControllerInterface(userService)

	router := gin.Default()
	routes.InitRoutes(&router.RouterGroup, userController)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
