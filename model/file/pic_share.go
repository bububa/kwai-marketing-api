package file

import "github.com/bububa/kwai-marketing-api/model"

// PicShareRequest 图片库推送 API Request
type PicShareRequest struct {
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// PicIDs 分享图片的ids
	// 分享图片的ids和tokens只能二选一
	PicIDs []string `json:"pic_ids,omitempty"`
	// ImageTokens 分享图片的tokens
	ImageTokens []string `json:"image_tokens,omitempty"`
	// ShareAdvertiserIDs 推送账户
	ShareAdvertiserIDs []uint64 `json:"share_advertiser_ids,omitempty"`
	// ShareAccountType 分享账户类型
	// 1：同主体同代理商账户；2：同主体下账户（需加白使用）； 3：内部代理商账户（需要加白使用）
	ShareAccountType int `json:"share_account_type,omitempty"`
}

// Url implement PostRequest interface
func (r PicShareRequest) Url() string {
	return "gw/dsp/v1/pic/share"
}

// Encode implement PostRequest interface
func (r PicShareRequest) Encode() []byte {
	return model.JSONMarshal(r)
}

// PicShareResponse 图片库推送 API Response
type PicShareResponse struct {
	// Details 详细信息
	Details []PicShareResult `json:"details,omitempty"`
}

// PicShareResult
type PicShareResult struct {
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// ImageToken 图片 token
	ImageToken string `json:"image_token,omitempty"`
	// PicID 图片id
	PicID string `json:"pic_id,omitempty"`
}
