package controller

import (
	"net/http"

	"github.com/corrots/cloud-storage/code"
	"github.com/corrots/cloud-storage/pkg/response"

	"github.com/gin-gonic/gin"
)

func (a *ApiCtrl) UploadHandler(c *gin.Context) {
	fh, err := c.FormFile("file")
	if err != nil {
		panic(code.ErrUploaded)
	}

	err = a.FileService.Save(fh)
	if err != nil {
		panic(code.ErrInternalServer)
	}

	c.JSON(http.StatusOK, response.OK(nil))
}
