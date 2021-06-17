package controller

import (
	"github.com/corrots/cloud-storage/code"
	"github.com/corrots/cloud-storage/pkg/xgin"
	"github.com/gin-gonic/gin"
)

func mustBindContext(ctx *gin.Context, req interface{}) {
	if err := xgin.ContextBindWithValid(ctx, req); err != nil {
		logger.Errorf("parameter validation err: %v\n", err)
		panic(code.ErrParameter())
	}
}
