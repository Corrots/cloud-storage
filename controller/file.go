package controller

import (
	"github.com/corrots/cloud-storage/code"
	"github.com/gin-gonic/gin"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"time"
)

func UploadHandler(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		panic(code.ErrInternalServer)
	}

	file := form.File["file"][0]
	dst := "tmp/" + time.Now().Format("20060102150405-") + file.Filename
	//fmt.Println("dst: ", dst)
	if err := saveUploadedFile(file, dst); err != nil {
		panic(code.ErrInternalServer)
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "ok",
	})
}

func saveUploadedFile(file *multipart.FileHeader, dst string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, src)
	return err
}
