package main

import (
	"github.com/corrots/cloud-storage/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()

	userAPI := router.Group("/user")
	{
		userAPI.GET("/login", handler.UserLogin)
	}

	router.Run(":8088")
}
