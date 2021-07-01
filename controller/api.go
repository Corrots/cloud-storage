package controller

import "github.com/corrots/cloud-storage/service"

type ApiCtrl struct {
	FileService *service.FileService
}

func New(fileService *service.FileService) *ApiCtrl {
	return &ApiCtrl{
		FileService: fileService,
	}
}
