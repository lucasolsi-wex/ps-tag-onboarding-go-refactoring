package controller

import "github.com/gin-gonic/gin"

func GetUserById(gc *gin.Context) {
	gc.JSON(200, "ok")
}
