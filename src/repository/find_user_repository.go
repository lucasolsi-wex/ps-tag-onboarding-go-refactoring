package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"lucasolsi-wex/ps-tag-onboarding-go/src/custom_errors"
	"lucasolsi-wex/ps-tag-onboarding-go/src/model"
	"lucasolsi-wex/ps-tag-onboarding-go/src/model/entity"
	"lucasolsi-wex/ps-tag-onboarding-go/src/utils"
)

func (userRepo *userRepository) FindUserById(id string) (model.UserDomainInterface, *custom_errors.CustomErr) {
	collectionName := viper.GetString(MongoDBUserDb)
	collection := userRepo.databaseConnection.Collection(collectionName)

	userEntity := &entity.UserEntity{}

	objectId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{
		Key:   "_id",
		Value: objectId,
	}}
	err := collection.FindOne(context.Background(), filter).Decode(userEntity)

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			errorMessage := fmt.Sprintf("User not found with id: %s", id)
			return nil, custom_errors.NewUserNotFoundError(errorMessage)
		}
		errorMessage := "Error in find user by id"
		return nil, custom_errors.NewInternalServerError(errorMessage)
	}

	return utils.ConvertEntityToDomain(*userEntity), nil
}
