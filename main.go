package main

import (
	"github.com/corrots/cloud-storage/controller"
	"github.com/corrots/cloud-storage/pkg/xgin"
)

func main() {
	router := xgin.New()

	userAPI := router.Group("/user")
	{
		userAPI.POST("/login", controller.UserLogin)
	}

	router.Run(":8088")
}
