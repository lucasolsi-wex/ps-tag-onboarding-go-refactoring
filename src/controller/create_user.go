package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"lucasolsi-wex/ps-tag-onboarding-go/src/model"
	"lucasolsi-wex/ps-tag-onboarding-go/src/model/request"
	"lucasolsi-wex/ps-tag-onboarding-go/src/utils"
	"net/http"
)

func (uc *userControllerInterface) CreateUser(gc *gin.Context) {
	var userRequest request.UserRequest

	if err := gc.ShouldBindJSON(userRequest); err != nil {
		customErr := utils.ValidateUserError(err)
		gc.JSON(customErr.Code, customErr)
		return
	}
	fmt.Println(userRequest)

	domain := model.NewUserDomain(userRequest.FirstName, userRequest.LastName, userRequest.Email, userRequest.Age)
	domainResult, err := uc.service.CreateUser(domain)

	if err != nil {
		gc.JSON(err.Code, err)
		return
	}

	gc.JSON(http.StatusCreated, utils.ConvertDomainToResponse(domainResult))
}
