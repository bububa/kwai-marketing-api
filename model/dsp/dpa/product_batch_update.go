package dpa

import "github.com/bububa/kwai-marketing-api/model"

// ProductBatchUpdateRequest 更新商品 API Request
type ProductBatchUpdateRequest struct {
	// DpaProductEditParams 批量创建商品数据
	// 单次请求最大数据量20
	DpaProductEditParams []ProductInfo `json:"dpa_product_edit_params,omitempty"`
	// AdvertiserID 广告主账号ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
}

// Url implements PostRequest interface
func (r ProductBatchUpdateRequest) Url() string {
	return "gw/dsp/v1/dpa/product/batch/update"
}

// Encode implements PostRequest interface
func (r ProductBatchUpdateRequest) Encode() []byte {
	return model.JSONMarshal(r)
}
