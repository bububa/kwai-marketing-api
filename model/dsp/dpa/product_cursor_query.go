package dpa

import "github.com/bububa/kwai-marketing-api/model"

// ProductCursorQueryRequest 获取商品列表(游标) API Request
type ProductCursorQueryRequest struct {
	// Cursor 游标	首次查询后返回，后续滚动查询必填，有效期3分钟
	Cursor string `json:"cursor,omitempty"`
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// LibraryID 商品库ID
	LibraryID uint64 `json:"library_id,omitempty"`
	// Limit 查询最大数据量	最大值500，无cursor时必填，有cursor时非必填且不生效
	Limit int `json:"limit,omitempty"`
}

// Url implement PostRequest interface
func (r ProductCursorQueryRequest) Url() string {
	return "gw/dsp/dpa/product/cursor_query"
}

// Encode implement PostRequest interface
func (r ProductCursorQueryRequest) Encode() []byte {
	return model.JSONMarshal(r)
}

// ProductCursorQueryResponse 获取商品列表(游标) API Response
type ProductCursorQueryResponse struct {
	// Cursor 游标
	Cursor string `json:"cursor,omitempty"`
	// ProductInfo 商品信息集
	ProductInfo []ProductInfo `json:"product_info,omitempty"`
}
