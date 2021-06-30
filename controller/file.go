package controller

import (
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"path/filepath"

	"github.com/corrots/cloud-storage/code"
	"github.com/corrots/cloud-storage/pkg/crypto"
	"github.com/corrots/cloud-storage/pkg/errors"
	"github.com/corrots/cloud-storage/pkg/files"
	"github.com/corrots/cloud-storage/pkg/response"

	"github.com/gin-gonic/gin"
)

const tempDir = "./tmp"

func UploadHandler(c *gin.Context) {
	fh, err := c.FormFile("file")
	if err != nil {
		panic(code.ErrUploaded)
	}

	if err := save(fh); err != nil {
		panic(code.ErrInternalServer)
	}

	c.JSON(http.StatusOK, response.New(nil))
}

func save(fh *multipart.FileHeader) error {
	file, err := fh.Open()
	if err != nil {
		return err
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	if err := files.MkdirAll("./tmp"); err != nil {
		return err
	}

	dst := getFilepath(crypto.ToMD5(bytes), fh.Filename)
	if err = files.SaveUploadedFile(fh, dst); err != nil {
		return errors.WithMessage(err, "save upload file err")
	}
	return nil
}

func getFilepath(hashed, name string) string {
	return filepath.Join(tempDir, hashed+"-"+name)
}
