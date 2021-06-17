package controller

import (
	"net/http"

	"github.com/corrots/cloud-storage/code"
	"github.com/corrots/cloud-storage/model"
	"github.com/corrots/cloud-storage/pkg/logging"
	"github.com/corrots/cloud-storage/pkg/response"

	"github.com/gin-gonic/gin"
)

var logger = logging.MustGetLogger("user")

func UserLogin(c *gin.Context) {
	var req model.RequestLogin

	mustBindContext(c, &req)

	logger.Info(req.Username)

	if req.Username != "admin" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, response.Err(code.ErrInternalServer))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "ok",
	})
}
