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

func (a *ApiCtrl) Download(c *gin.Context) {
	c.HTML(http.StatusOK, "download.html", nil)
}

func (a *ApiCtrl) SignIn(c *gin.Context) {
	c.HTML(http.StatusOK, "signin.html", nil)
}

func (a *ApiCtrl) SignUp(c *gin.Context) {
	c.HTML(http.StatusOK, "signup.html", nil)
}
