package utils

import (
	"lucasolsi-wex/ps-tag-onboarding-go/src/model"
	"lucasolsi-wex/ps-tag-onboarding-go/src/model/response"
)

func ConvertDomainToResponse(domainInterface model.UserDomainInterface) response.UserResponse {
	return response.UserResponse{
		Id:        domainInterface.GetId(),
		FirstName: domainInterface.GetFirstName(),
		LastName:  domainInterface.GetLastName(),
		Email:     domainInterface.GetEmail(),
		Age:       domainInterface.GetAge(),
	}
}
