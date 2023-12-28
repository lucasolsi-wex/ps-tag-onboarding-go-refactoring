package repository

import (
	"go.mongodb.org/mongo-driver/mongo"
	"lucasolsi-wex/ps-tag-onboarding-go/src/custom_errors"
	"lucasolsi-wex/ps-tag-onboarding-go/src/service"
)

func NewUserRepository(database *mongo.Database) UserRepository {
	return &userRepository{databaseConnection: database}
}

type userRepository struct {
	databaseConnection *mongo.Database
}

type UserRepository interface {
	CreateUser(domainInterface service.UserDomainInterface) (service.UserDomainInterface, *custom_errors.CustomErr)
}
