package controller

import (
	"net/http"

	"github.com/corrots/cloud-storage/code"
	"github.com/corrots/cloud-storage/pkg/errors"
	"github.com/corrots/cloud-storage/pkg/logging"
	"github.com/corrots/cloud-storage/pkg/xgin"
	"github.com/gin-gonic/gin"
	"github.com/isnlan/coral/pkg/response"
)

var logger = logging.MustGetLogger("controller")

func mustBindContext(ctx *gin.Context, req interface{}) {
	if err := xgin.ContextBindWithValid(ctx, req); err != nil {
		logger.Errorf("parameter validation err: %v\n", err)
		panic(code.ErrParameter())
	}
}

func mustBindQuery(ctx *gin.Context, req interface{}) {
	if err := xgin.ContextBindQueryWithValid(ctx, req); err != nil {
		logger.Errorf("parameter validation err: %v\n", err)
		panic(code.ErrParameter())
	}
}

func (a *ApiCtrl) response(ctx *gin.Context, data interface{}) {
	switch data.(type) {
	default:
		ctx.JSON(http.StatusOK, response.New(data))
	}
}

func (a *ApiCtrl) html(ctx *gin.Context, name string) {
	ctx.HTML(http.StatusOK, name, nil)
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
