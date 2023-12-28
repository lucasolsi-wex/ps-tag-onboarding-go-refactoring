package controller

import (
	"github.com/gin-gonic/gin"
	"lucasolsi-wex/ps-tag-onboarding-go/src/service"
)

func NewUserControllerInterface(domainService service.UserDomainService) UserControllerInterface {
	return &userControllerInterface{service: domainService}
}

type UserControllerInterface interface {
	CreateUser(gc *gin.Context)
	FindUserById(gc *gin.Context)
}

type userControllerInterface struct {
	service service.UserDomainService
}
