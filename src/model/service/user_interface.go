package service

import (
	"lucasolsi-wex/ps-tag-onboarding-go/src/custom_errors"
	"lucasolsi-wex/ps-tag-onboarding-go/src/model"
	"lucasolsi-wex/ps-tag-onboarding-go/src/repository"
)

func NewUserDomainService(repository repository.UserRepository) UserDomainService {
	return &userDomainService{repository}
}

type userDomainService struct {
	repository repository.UserRepository
}

type UserDomainService interface {
	CreateUser(domainInterface model.UserDomainInterface) (model.UserDomainInterface, *custom_errors.CustomErr)
	FindUserById(string) (domainInterface model.UserDomainInterface, err *custom_errors.CustomErr)
}
