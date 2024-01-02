package service

import (
	"lucasolsi-wex/ps-tag-onboarding-go/src/custom_errors"
	"lucasolsi-wex/ps-tag-onboarding-go/src/model"
)

func (ud *userDomainService) CreateUser(
	domain model.UserDomainInterface) (model.UserDomainInterface, *custom_errors.CustomErr) {

	userDomainRepo, err := ud.repository.CreateUser(domain)
	if err != nil {
		return nil, err
	}

	return userDomainRepo, err
}
