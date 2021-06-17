package xgin

import (
	"github.com/corrots/cloud-storage/pkg/logging"
	"github.com/corrots/cloud-storage/pkg/trace"
	"github.com/gin-gonic/gin"
)

var logger = logging.MustGetLogger("xgin")

func New(middleware ...gin.HandlerFunc) *gin.Engine {
	middleware = append(middleware, LoggerWriter(), RecoveryWriter(), CorsMiddleware(), trace.TracerWrapper)

	router := gin.New()
	router.Use(middleware...)
	router.NoRoute(HandleNotFound)
	router.NoMethod(HandleNotFound)
	return router
}
