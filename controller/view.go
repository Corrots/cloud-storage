package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (a *ApiCtrl) Home(c *gin.Context) {
	c.HTML(http.StatusOK, "home.html", nil)
}

func (a *ApiCtrl) Upload(c *gin.Context) {
	c.HTML(http.StatusOK, "upload.html", nil)
}
