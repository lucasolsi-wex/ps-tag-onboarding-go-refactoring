package repository

import (
	"context"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"lucasolsi-wex/ps-tag-onboarding-go/src/custom_errors"
	"lucasolsi-wex/ps-tag-onboarding-go/src/service"
	"lucasolsi-wex/ps-tag-onboarding-go/src/utils"
)

const (
	MongoDBUserDb = "MONGODB_DATABASE_COLLECTION"
)

func (repo *userRepository) CreateUser(domainInterface service.UserDomainInterface) (
	service.UserDomainInterface,
	*custom_errors.CustomErr) {
	collectionName := viper.GetString(MongoDBUserDb)
	collection := repo.databaseConnection.Collection(collectionName)

	value := utils.ConvertDomainToEntity(domainInterface)

	result, err := collection.InsertOne(context.Background(), value)

	if err != nil {
		return nil, custom_errors.NewInternalServerError(err.Error())
	}

	value.Id = result.InsertedID.(primitive.ObjectID)

	return utils.ConvertEntityToDomain(*value), nil

}
