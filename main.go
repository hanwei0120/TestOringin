package main

import (
	"gitTest2/controller"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	router := gin.Default()
	router.GET("/stuInfo",controller.GetInfo)
	router.POST("/bookInfo",controller.GetBookInfo)
	if err := router.Run("127.0.0.1:9797"); err != nil {
		log.Fatal("router.Run")
	}
}



