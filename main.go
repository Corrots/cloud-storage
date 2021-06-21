package main

import (
	"net/http"

	"github.com/corrots/cloud-storage/controller"
	"github.com/corrots/cloud-storage/pkg/xgin"
	"github.com/gin-gonic/gin"
)

func main() {
	router := xgin.New()

	HomeAPI := router.Group("/index")
	{
		router.LoadHTMLGlob("static/view/*")
		router.Static("/static", "./static")
		HomeAPI.GET("/", func(c *gin.Context) {
			c.HTML(http.StatusOK, "home.html", nil)
		})
	}

	fileAPI := router.Group("/file")
	{

		fileAPI.GET("/query")
		fileAPI.GET("/download")
		fileAPI.POST("/upload")
		fileAPI.POST("/update")
		fileAPI.POST("/delete")
	}

	userAPI := router.Group("/user")
	{
		userAPI.POST("/login", controller.UserLogin)
	}

	router.Run(":8080")
}
