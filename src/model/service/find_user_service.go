package service

import (
	"lucasolsi-wex/ps-tag-onboarding-go/src/custom_errors"
	"lucasolsi-wex/ps-tag-onboarding-go/src/model"
)

func (ud *userDomainService) FindUserById(id string) (model.UserDomainInterface, *custom_errors.CustomErr) {
	return ud.repository.FindUserById(id)
}
