package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
	"lucasolsi-wex/ps-tag-onboarding-go/src/config/database"
	"lucasolsi-wex/ps-tag-onboarding-go/src/controller/routes"
)

func main() {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		return
	}

	dbConnection, err := database.NewMongoDBConnection(context.Background())

	userController := initDependencies(dbConnection)

	router := gin.Default()
	routes.InitRoutes(&router.RouterGroup, userController)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
