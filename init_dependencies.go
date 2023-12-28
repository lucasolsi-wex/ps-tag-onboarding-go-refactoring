package main

import (
	"go.mongodb.org/mongo-driver/mongo"
	"lucasolsi-wex/ps-tag-onboarding-go/src/controller"
	"lucasolsi-wex/ps-tag-onboarding-go/src/model/service"
	"lucasolsi-wex/ps-tag-onboarding-go/src/repository"
)

func initDependencies(database *mongo.Database) controller.UserControllerInterface {
	repo := repository.NewUserRepository(database)
	domainService := service.NewUserDomainService(repo)
	return controller.NewUserControllerInterface(domainService)
}
