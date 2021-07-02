package controller

import (
	"github.com/gin-gonic/gin"
)

func (a *ApiCtrl) Home(c *gin.Context) {
	a.html(c, "home.html")
}

func (a *ApiCtrl) Upload(c *gin.Context) {
	a.html(c, "upload.html")
}

func (a *ApiCtrl) Download(c *gin.Context) {
	a.html(c, "download.html")
}

func (a *ApiCtrl) SignIn(c *gin.Context) {
	a.html(c, "signin.html")
}

func (a *ApiCtrl) SignUp(c *gin.Context) {
	a.html(c, "signup.html")
}
