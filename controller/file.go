package controller

import (
	"github.com/corrots/cloud-storage/code"
	"github.com/gin-gonic/gin"
)

func (a *ApiCtrl) FileUpload(ctx *gin.Context) {
	fh, err := ctx.FormFile("file")
	if err != nil {
		panic(code.ErrUploaded)
	}

	err = a.FileService.Save(fh)
	if err != nil {
		logger.Errorf("file save err: %v\n", err)
		panic(code.ErrInternalServer)
	}

	a.response(ctx, nil)
}

func (a *ApiCtrl) FileDownload(ctx *gin.Context) {
	filename := "b56a11515cb1062ad5716069ccc0931b-p2620309098.jpg"
	path, err := a.FileService.Download(filename)
	if err != nil {
		logger.Errorf("file download err: %v\n", err)
		panic(code.ErrInternalServer)
	}
	ctx.FileAttachment(path, filename)
}
