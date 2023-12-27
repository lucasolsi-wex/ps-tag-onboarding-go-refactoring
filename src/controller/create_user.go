package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"lucasolsi-wex/ps-tag-onboarding-go/src/custom_errors"
	"lucasolsi-wex/ps-tag-onboarding-go/src/model/request"
)

func CreateUser(gc *gin.Context) {
	var userRequest request.UserRequest

	if err := gc.ShouldBindJSON(userRequest); err != nil {
		customErr := custom_errors.NewBadRequestError(fmt.Sprintf("Something went wrong with the following fields: %s\n",
			err.Error()))
		gc.JSON(customErr.Code, customErr)
		return
	}
	fmt.Println(userRequest)
}
