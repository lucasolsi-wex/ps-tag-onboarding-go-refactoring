package service

import "lucasolsi-wex/ps-tag-onboarding-go/src/custom_errors"

func NewUserDomainService() UserDomainService {
	return &userDomainService{}
}

type userDomainService struct {
}

type UserDomainService interface {
	CreateUser(domainInterface UserDomainInterface) *custom_errors.CustomErr
	FindUser(string) (domainInterface *UserDomainInterface, err *custom_errors.CustomErr)
}
