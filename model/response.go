package model

import "encoding/json"

// Response api response interface
type Response interface {
	IsError() bool
	Error() string
}

// BaseResponse shared response data struct
type BaseResponse struct {
	Message   string          `json:"message,omitempty"`    // 返回信息
	RequestId string          `json:"request_id,omitempty"` // 请求id
	Data      json.RawMessage `json:"data,omitempty"`       // JSON返回值
	Code      int             `json:"code,omitempty"`       // 返回码
}

// IsError detect if the response is an error
func (r BaseResponse) IsError() bool {
	return r.Code != 0
}

// Error implement error interface
func (r BaseResponse) Error() string {
	return r.Message
}
