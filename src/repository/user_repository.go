package repository

import (
	"go.mongodb.org/mongo-driver/mongo"
	"lucasolsi-wex/ps-tag-onboarding-go/src/custom_errors"
	"lucasolsi-wex/ps-tag-onboarding-go/src/model"
)

func NewUserRepository(database *mongo.Database) UserRepository {
	return &userRepository{databaseConnection: database}
}

type userRepository struct {
	databaseConnection *mongo.Database
}

type UserRepository interface {
	CreateUser(domainInterface model.UserDomainInterface) (model.UserDomainInterface, *custom_errors.CustomErr)
	FindUserById(id string) (model.UserDomainInterface, *custom_errors.CustomErr)
	ExistsByFirstNameAndLastName(firstName, lastName string) bool
}
