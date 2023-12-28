package utils

import (
	"lucasolsi-wex/ps-tag-onboarding-go/src/model/entity"
	"lucasolsi-wex/ps-tag-onboarding-go/src/service"
)

func ConvertEntityToDomain(entity entity.UserEntity) service.UserDomainInterface {
	domain := service.NewUserDomain(entity.FirstName, entity.LastName, entity.Email, entity.Age)
	domain.SetId(entity.Id.Hex())

	return domain
}
