package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//页面显示111222333444555，未调用grpc
func GetInfo(ctx *gin.Context) {
	ctx.String(http.StatusOK, "111222333444555")
	ctx.JSON(200,gin.H{
		"what":"111222333444555",
	})
}
