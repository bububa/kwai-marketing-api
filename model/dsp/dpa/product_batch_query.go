package dpa

import "github.com/bububa/kwai-marketing-api/model"

// ProductBatchQueryRequest 获取商品列表 API Request
type ProductBatchQueryRequest struct {
	// AdDpaProductBatchQueryParamSneak 商品批量查询条件
	AdDpaProductBatchQueryParamSneak *AdDpaProductBatchQueryParamSneak `json:"dpa_product_query_param,omitempty"`
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
}

// Url implement PostRequest interface
func (r ProductBatchQueryRequest) Url() string {
	return "gw/dsp/v1/dpa/product/batch/query"
}

// Encode implement PostRequest interface
func (r ProductBatchQueryRequest) Encode() []byte {
	return model.JSONMarshal(r)
}

// AdDpaProductBatchQueryParamSneak 商品批量查询条件
type AdDpaProductBatchQueryParamSneak struct {
	// Name 商品名称
	Name string `json:"name,omitempty"`
	// 	商品第三方IDs
	OuterIDs []string `json:"outer_ids,omitempty"`
	// LibraryID 商品库ID
	LibraryID uint64 `json:"library_id,omitempty"`
	// Status 商品投放状态 0-不可投放, 1-可投放
	Status int `json:"status,omitempty"`
	// CheckStatus 商品接入(校验)状态 0-成功, 1-失败
	CheckStatus int `json:"check_status,omitempty"`
}

// ProductBatchQueryResponse 获取商品库列表 API Response
type ProductBatchQueryResponse struct {
	// PageInfo 列表页参数
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
	// ProductInfo 商品信息
	ProductInfo []ProductInfo `json:"product_info,omitempty"`
}
