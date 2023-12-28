package utils

import (
	"lucasolsi-wex/ps-tag-onboarding-go/src/model/entity"
	"lucasolsi-wex/ps-tag-onboarding-go/src/service"
)

func ConvertDomainToEntity(domain service.UserDomainInterface) *entity.UserEntity {

	return &entity.UserEntity{
		FirstName: domain.GetFirstName(),
		LastName:  domain.GetLastName(),
		Email:     domain.GetEmail(),
		Age:       domain.GetAge(),
	}
}
