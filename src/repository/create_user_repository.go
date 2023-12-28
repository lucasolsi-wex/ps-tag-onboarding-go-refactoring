package repository

import (
	"context"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"lucasolsi-wex/ps-tag-onboarding-go/src/custom_errors"
	"lucasolsi-wex/ps-tag-onboarding-go/src/model"
	"lucasolsi-wex/ps-tag-onboarding-go/src/utils"
)

const (
	MongoDBUserDb = "MONGODB_DATABASE_COLLECTION"
)

func (userRepo *userRepository) CreateUser(domainInterface model.UserDomainInterface) (
	model.UserDomainInterface,
	*custom_errors.CustomErr) {
	collectionName := viper.GetString(MongoDBUserDb)
	collection := userRepo.databaseConnection.Collection(collectionName)

	value := utils.ConvertDomainToEntity(domainInterface)

	result, err := collection.InsertOne(context.Background(), value)

	if err != nil {
		return nil, custom_errors.NewInternalServerError(err.Error())
	}

	value.Id = result.InsertedID.(primitive.ObjectID)

	return utils.ConvertEntityToDomain(*value), nil
}
