package utils

import (
	"lucasolsi-wex/ps-tag-onboarding-go/src/model/response"
	"lucasolsi-wex/ps-tag-onboarding-go/src/service"
)

func ConvertDomainToResponse(domainInterface service.UserDomainInterface) response.UserResponse {
	return response.UserResponse{
		Id:        domainInterface.GetId(),
		FirstName: domainInterface.GetFirstName(),
		LastName:  domainInterface.GetLastName(),
		Email:     domainInterface.GetEmail(),
		Age:       domainInterface.GetAge(),
	}
}
