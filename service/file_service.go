package service

import (
	"io/ioutil"
	"mime/multipart"
	"path/filepath"

	"github.com/corrots/cloud-storage/pkg/crypto"
	"github.com/corrots/cloud-storage/pkg/errors"
	"github.com/corrots/cloud-storage/pkg/files"
	"github.com/corrots/cloud-storage/pkg/logging"
)

type FileService struct {
}

var logger = logging.MustGetLogger("service")

func NewFileService() *FileService {
	return &FileService{}
}

func (svc *FileService) Save(fh *multipart.FileHeader) error {
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

const tempDir = "./tmp"

func getFilepath(hashed, name string) string {
	return filepath.Join(tempDir, hashed+"-"+name)
}
