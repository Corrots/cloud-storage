package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	c.HTML(http.StatusOK, "home.html", nil)
}

func Upload(c *gin.Context) {
	c.HTML(http.StatusOK, "upload.html", nil)
}
