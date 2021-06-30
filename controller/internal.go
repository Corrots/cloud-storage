package controller

import (
	"github.com/corrots/cloud-storage/code"
	"github.com/corrots/cloud-storage/pkg/errors"
	"github.com/corrots/cloud-storage/pkg/logging"
	"github.com/corrots/cloud-storage/pkg/xgin"
	"github.com/gin-gonic/gin"
)

var logger = logging.MustGetLogger("controller")

func mustBindContext(ctx *gin.Context, req interface{}) {
	if err := xgin.ContextBindWithValid(ctx, req); err != nil {
		logger.Errorf("parameter validation err: %v\n", err)
		panic(code.ErrParameter())
	}
}

func checkError(err error) {
	if err != nil {
		switch e := err.(type) {
		case errors.CodeError:
			panic(e)
		default:
			logger.Error(e)
			panic(code.ErrInternalServer)
		}
	}
}
