package controller

import (
	"net/http"

	"github.com/corrots/cloud-storage/code"
	"github.com/corrots/cloud-storage/model"
	"github.com/corrots/cloud-storage/pkg/response"

	"github.com/gin-gonic/gin"
)

func (a *ApiCtrl) UserLogin(c *gin.Context) {
	var req model.RequestLogin

	mustBindContext(c, &req)

	if req.Username != "admin" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, response.Err(code.ErrInternalServer))
		return
	}

	c.JSON(http.StatusOK, response.OK(nil))
}
