package http

import (
	"fmt"

	"github.com/corrots/cloud-storage/controller"
	"github.com/corrots/cloud-storage/pkg/xgin"

	"github.com/gin-gonic/gin"
)

type Server struct {
	addr   string
	router *gin.Engine
}

func New(addr string, ctrl *controller.ApiCtrl) *Server {
	router := xgin.New()

	fileAPI := router.Group("/file")
	{
		fileAPI.GET("/query")
		fileAPI.GET("/download", ctrl.FileDownload)
		fileAPI.POST("/upload", ctrl.FileUpload)
		fileAPI.POST("/update")
		fileAPI.POST("/delete")
	}

	userAPI := router.Group("/user")
	{
		userAPI.POST("/login", ctrl.UserLogin)
	}

	View := router.Group("/")
	{
		router.LoadHTMLGlob("static/view/*")
		router.Static("/static", "./static")
		View.GET("/", ctrl.Home)
		View.GET("/upload", ctrl.Upload)
		View.GET("/download", ctrl.Download)
		View.GET("/signin", ctrl.SignIn)
		View.GET("/signup", ctrl.SignUp)
	}

	return &Server{
		addr:   addr,
		router: router,
	}
}

func (s *Server) Start() {
	err := s.router.Run(s.addr)
	if err != nil {
		panic(fmt.Sprintf("start http server: [%s] err: %v\n", s.addr, err))
	}
}
