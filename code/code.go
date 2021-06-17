package code

import "github.com/corrots/cloud-storage/pkg/errors"

var (
	ErrInternalServer = errors.NewWithInfo(500, "服务器发生错误，请稍后再试")
)
