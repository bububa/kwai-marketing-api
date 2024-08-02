package dpa

import (
	"fmt"

	"github.com/bububa/kwai-marketing-api/model"
)

// CreativeVideoGenerateRequest 批量模板合成SDPA创意视频 API Request
type CreativeVideoGenerateRequest struct {
	// OuterID 商品外部ID
	OuterID string `json:"outer_id,omitempty"`
	// Templates 视频模板信息
	Templates []CreativeTemplate `json:"templates,omitempty"`
	// AdvertiserID 广告主账号ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// LibraryID 商品库ID
	LibraryID uint64 `json:"library_id,omitempty"`
	// ProductID 商品ID
	// 优先于"outer_id"生效
	ProductID uint64 `json:"product_id,omitempty"`
}

// Url implements PostRequest interface
func (r CreativeVideoGenerateRequest) Url() string {
	return "gw/dsp/v1/dpa/creative/video/generate"
}

// Encode implements PostRequest interface
func (r CreativeVideoGenerateRequest) Encode() []byte {
	return model.JSONMarshal(r)
}

// CreativeVideoGenerateResponse 批量模板合成SDPA创意视频 API Response
type CreativeVideoGenerateResponse struct {
	// VideoInfos 视频信息列表
	VideoInfos []GenerateVideoResult `json:"video_infos,omitempty"`
}

// GenerateVideoResult 合成视频信息
type GenerateVideoResult struct {
	// OuterID 商品第三方ID
	OuterID string `json:"outer_id,omitempty"`
	// ProductName 商品名称
	ProductName string `json:"product_name,omitempty"`
	// VideoURL 视频URL
	VideoURL string `json:"video_url,omitempty"`
	// Code 失败状态码
	Code string `json:"code,omitempty"`
	// Message 失败信息
	Message string `json:"message,omitempty"`
	// PhotoID 	Photo ID
	PhotoID uint64 `json:"photo_id,omitempty"`
	// TemplateID 模板ID
	TemplateID uint64 `json:"template_id,omitempty"`
	// ProductID 商品ID
	ProductID uint64 `json:"product_id,omitempty"`
}

// IsError check generated video is failed
func (r GenerateVideoResult) IsError() bool {
	return r.Code != ""
}

// Error implements error interface
func (r GenerateVideoResult) Error() string {
	return fmt.Sprintf("code:%s, msg:%s", r.Code, r.Message)
}
