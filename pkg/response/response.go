package response

import "github.com/corrots/cloud-storage/pkg/errors"

const SuccessCode = 0

const SuccessMsg = "ok"

type Response struct {
	ErrorCode   int         `json:"errcode"`
	Description string      `json:"description"`
	Data        interface{} `json:"data,omitempty"`
}

func OK(data interface{}) *Response {
	return &Response{
		ErrorCode:   SuccessCode,
		Description: SuccessMsg,
		Data:        data,
	}
}

func Fail(code int, desc string) *Response {
	return &Response{
		ErrorCode:   code,
		Description: desc,
		Data:        nil,
	}
}

func Err(err error) *Response {
	if e, ok := err.(errors.CodeError); ok {
		return &Response{
			ErrorCode:   e.Code(),
			Description: e.Error(),
			Data:        nil,
		}
	}

	return &Response{
		ErrorCode:   errors.InternalErrorCode,
		Description: err.Error(),
	}
}
