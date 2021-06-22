package main

import (
	"github.com/corrots/cloud-storage/controller"
	"github.com/corrots/cloud-storage/pkg/xgin"
)

func main() {
	router := xgin.New()

	View := router.Group("/")
	{
		router.LoadHTMLGlob("static/view/*")
		router.Static("/static", "./static")
		View.GET("/", controller.Home)
		View.GET("/upload", controller.Upload)
	}

	fileAPI := router.Group("/file")
	{
		fileAPI.GET("/query")
		fileAPI.GET("/download")
		fileAPI.POST("/upload", controller.UploadHandler)
		fileAPI.POST("/update")
		fileAPI.POST("/delete")
	}

	userAPI := router.Group("/user")
	{
		userAPI.POST("/login", controller.UserLogin)
	}

	router.Run(":8080")
}
