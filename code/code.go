package code

import "github.com/corrots/cloud-storage/pkg/errors"

var (
	ErrInternalServer = errors.NewWithInfo(500, "服务器发生错误，请稍后再试")
)

var ErrParameter = func(info ...string) errors.CodeError {
	msg := "请求参数错误"
	if len(info) > 0 {
		msg = info[0]
	}
	return errors.NewWithInfo(6000, msg)
}
