package dpa

import "github.com/bububa/kwai-marketing-api/model"

// ProductBatchCreateRequest 创建商品 API Request
type ProductBatchCreateRequest struct {
	// DpaProductEditParams 批量创建商品数据
	// 单次请求最大数据量20
	DpaProductEditParams []ProductInfo `json:"dpa_product_edit_params,omitempty"`
	// AdvertiserID 广告主账号ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
}

// Url implements PostRequest interface
func (r ProductBatchCreateRequest) Url() string {
	return "gw/dsp/v1/dpa/product/batch/create"
}

// Encode implements PostRequest interface
func (r ProductBatchCreateRequest) Encode() []byte {
	return model.JSONMarshal(r)
}

// ProductBatchUpdateResponse 创建/编辑商品 API Response
type ProductBatchUpdateResponse struct {
	ProductEditResponses []ProductUpdateResult `json:"product_edit_responses,omitempty"`
}

type ProductUpdateResult struct {
	// OuterID 商品外部ID
	OuterID string `json:"outer_id,omitempty"`
	// Success 创建/编辑结果
	// true-成功, false-失败
	// ErrorMsg 错误信息
	ErrorMsg string `json:"error_msg,omitempty"`
	// LibraryID 商品库ID
	LibraryID uint64 `json:"library_id,omitempty"`
	// ProductID 商品ID
	// 快手生成商品唯一标识
	ProductID uint64 `json:"product_id,omitempty"`
}
