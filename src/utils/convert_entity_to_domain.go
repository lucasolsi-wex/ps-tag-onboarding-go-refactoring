package utils

import (
	"lucasolsi-wex/ps-tag-onboarding-go/src/model"
	"lucasolsi-wex/ps-tag-onboarding-go/src/model/entity"
)

func ConvertEntityToDomain(entity entity.UserEntity) model.UserDomainInterface {
	domain := model.NewUserDomain(entity.FirstName, entity.LastName, entity.Email, entity.Age)
	domain.SetId(entity.Id.Hex())

	return domain
}
