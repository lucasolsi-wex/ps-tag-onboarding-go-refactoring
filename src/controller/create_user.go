package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"lucasolsi-wex/ps-tag-onboarding-go/src/model/request"
	"lucasolsi-wex/ps-tag-onboarding-go/src/service"
	"lucasolsi-wex/ps-tag-onboarding-go/src/utils"
	"lucasolsi-wex/ps-tag-onboarding-go/src/view"
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

	domain := service.NewUserDomain(userRequest.FirstName, userRequest.LastName, userRequest.Email, userRequest.Age)

	if err := uc.service.CreateUser(domain); err != nil {
		gc.JSON(err.Code, err)
		return
	}

	gc.JSON(http.StatusCreated, view.ConvertDomainToResponse(domain))
}
