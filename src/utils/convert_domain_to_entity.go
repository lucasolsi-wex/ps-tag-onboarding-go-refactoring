package utils

import (
	"lucasolsi-wex/ps-tag-onboarding-go/src/model"
	"lucasolsi-wex/ps-tag-onboarding-go/src/model/entity"
)

func ConvertDomainToEntity(domain model.UserDomainInterface) *entity.UserEntity {

	return &entity.UserEntity{
		FirstName: domain.GetFirstName(),
		LastName:  domain.GetLastName(),
		Email:     domain.GetEmail(),
		Age:       domain.GetAge(),
	}
}
