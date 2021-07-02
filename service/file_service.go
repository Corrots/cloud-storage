package service

import (
	"io/ioutil"
	"mime/multipart"
	"path/filepath"

	"github.com/corrots/cloud-storage/pkg/db"

	"github.com/jmoiron/sqlx"

	"github.com/corrots/cloud-storage/config"
	"github.com/corrots/cloud-storage/pkg/crypto"
	"github.com/corrots/cloud-storage/pkg/errors"
	"github.com/corrots/cloud-storage/pkg/files"
	"github.com/corrots/cloud-storage/pkg/logging"

	"github.com/mitchellh/go-homedir"
)

type FileService struct {
	db *sqlx.DB
}

var logger = logging.MustGetLogger("service")

func NewFileService() *FileService {
	return &FileService{
		db: db.GetSqlxDB(),
	}
}

func (svc *FileService) create() {

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

	if err := files.Mkdir(tempDir()); err != nil {
		return err
	}

	dst := fullFilepath(crypto.ToMD5(bytes), fh.Filename)
	if err = files.SaveUploadedFile(fh, dst); err != nil {
		return errors.WithMessage(err, "save upload file err")
	}
	return nil
}

func (svc *FileService) Download(filename string) (string, error) {
	path := filepath.Join(tempDir(), filename)
	ok, err := files.Exists(path)
	if err != nil {
		return "", err
	}
	if !ok {
		return "", errors.WithMessage(err, "file not exist")
	}
	return path, nil
}

func tempDir() string {
	// get the home directory for the executing user.
	home, err := homedir.Dir()
	if err != nil {
		panic(err)
	}
	return filepath.Join(home, config.GlobalConfig.Server.Tmpdir)
}

func fullFilepath(hashed, name string) string {
	return filepath.Join(tempDir(), hashed+"-"+name)
}
