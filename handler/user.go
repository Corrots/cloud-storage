package handler

import (
	"net/http"

	"github.com/corrots/cloud-storage/code"
	"github.com/corrots/cloud-storage/model"
	"github.com/corrots/cloud-storage/pkg/response"

	"github.com/gin-gonic/gin"
)

func UserLogin(c *gin.Context) {
	var req model.RequestLogin

	if err := c.ShouldBind(&req); err != nil {
		panic(err)
	}

	if req.Username != "admin" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, response.Err(code.ErrInternalServer))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "ok",
	})
}
