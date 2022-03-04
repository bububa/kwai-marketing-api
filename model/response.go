package model

import "encoding/json"

// Response api response interface
type Response interface {
	IsError() bool
	Error() string
}

// BaseResponse shared response data struct
type BaseResponse struct {
	Code      int             `json:"code,omitempty"`    // 返回码
	Message   string          `json:"message,omitempty"` // 返回信息
	Data      json.RawMessage `json:"data,omitempty"`    // JSON返回值
	RequestId int64           `json:"request_id"`		 // 请求id
}

// IsError detect if the response is an error
func (r BaseResponse) IsError() bool {
	return r.Code != 0
}

// Error implement error interface
func (r BaseResponse) Error() string {
	return r.Message
}
