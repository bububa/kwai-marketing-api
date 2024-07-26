package page

import (
	"encoding/json"

	"github.com/bububa/kwai-marketing-api/model"
)

// CidInfoUpdateRequest 魔力建站落地页更新CID信息 API Request
type CidInfoUpdateRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// PageID 查询落地页列表接口获取
	PageID string `json:"page_id,omitempty"`
	// PlatFormType 更新链接的电商平台类型：PDD：拼多多, WECHAT_MALL:微信商城，TAO_BAO：淘宝，tmall：天猫，JD：京东
	PlatFormType string `json:"plat_form_type,omitempty"`
	// DeeplinkURL 更新落地页中按钮组件 — 跳转应用 — 跳转链接如：taobao://huodong.m.taobao.com/act/snipcode.htm
	DeeplinkURL string `json:"deeplink_url,omitempty"`
	// FallbackH5URL 更新落地页中按钮组件 — 跳转应用 — H5链接如：https://www.taobao.com
	FallbackH5URL string `json:"fallback_h5_url,omitempty"`
}

// Url implement PostRequest interface
func (r CidInfoUpdateRequest) Url() string {
	return "gw/magicsite/v1/page/cid/info/update"
}

// Encode implement PostRequest interface
func (r CidInfoUpdateRequest) Encode() []byte {
	ret, _ := json.Marshal(r)
	return ret
}

// CidInfoUpdateResponse 魔力建站落地页更新CID信息 API Response
type CidInfoUpdateResponse struct {
	// PageID 更新成功的落地页ID
	PageID model.Uint64 `json:"page_id,omitempty"`
}
