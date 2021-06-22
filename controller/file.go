package controller

import (
	"github.com/gin-gonic/gin"
)

func UploadHandler(c *gin.Context) {
	c.FormFile()
}
