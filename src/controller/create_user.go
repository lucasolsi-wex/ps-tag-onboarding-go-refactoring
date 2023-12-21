package controller

import "github.com/gin-gonic/gin"

func CreateUser(gc *gin.Context) {
	gc.JSON(200, "ok")
}
