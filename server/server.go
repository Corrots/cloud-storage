package server

import (
	"github.com/corrots/cloud-storage/config"
	"github.com/corrots/cloud-storage/controller"
	"github.com/corrots/cloud-storage/http"
	"github.com/corrots/cloud-storage/pkg/application"
	"github.com/corrots/cloud-storage/pkg/db"
	"github.com/corrots/cloud-storage/pkg/errors"
	"github.com/corrots/cloud-storage/service"
)

type Server struct {
	http *http.Server
}

func New() *Server {
	s := new(Server)
	return s
}

func (s *Server) Initialize() error {
	// init config
	if err := application.InitConfig(&config.GlobalConfig); err != nil {
		return errors.WithMessage(err, "init config err")
	}
	// init redis

	// init mysql
	mysql := config.GlobalConfig.Dao.Mysql
	err := db.InitSqlxDB(mysql.Uri, mysql.Username, mysql.Password, mysql.DB)
	if err != nil {
		return errors.WithMessage(err, "init mysql err")
	}

	ctrl := controller.New(service.NewFileService())
	s.http = http.New(config.GlobalConfig.Server.Addr, ctrl)
	return nil
}

func (s *Server) Start() {
	go s.http.Start()
}

func (s *Server) Close() {
	// do noting
}
