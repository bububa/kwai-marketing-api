package model

import "encoding/json"

type Response interface {
	IsError() bool
	Error() string
}

type BaseResponse struct {
	Code    int             `json:"code,omitempty"`    // 返回码
	Message string          `json:"message,omitempty"` // 返回信息
	Data    json.RawMessage `json:"data,omitempty"`    // JSON返回值
}

func (r BaseResponse) IsError() bool {
	return r.Code != 0
}

func (r BaseResponse) Error() string {
	return r.Message
}
