package file

import (
	"github.com/bububa/kwai-marketing-api/model"
)

// AdVideoShareNewRequest 视频库-推送视频(新版) API Request
type AdVideoShareNewRequest struct {
	// AdvertiserID 广告主 id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// PhotoIDs 视频 ids，不超过 10 个
	PhotoIDs []string `json:"photo_ids,omitempty"`
	// AccountIDs 目标推送账户
	AccountIDs []uint64 `json:"account_ids,omitempty"`
	// ShareAccountType 分享账户类型：-1：暂未支持的账户类型；1. 同主体同代理商账户；2. 同主体下账户（需加白使用）；3. 内部账户；4. 同代理商下账户
	ShareAccountType int `json:"share_account_type,omitempty"`
}

// Url implement PostRequest interface
func (r AdVideoShareNewRequest) Url() string {
	return "gw/dsp/v1/file/ad/video/share/new"
}

// Encode implement PostRequest interface
func (r AdVideoShareNewRequest) Encode() []byte {
	return model.JSONMarshal(r)
}

// AdVideoShareNewResponse 视频库-推送视频(新版) API Response
type AdVideoShareNewResponse struct {
	// ShareStatus 视频推送状态：-1=全部失败，0=表示部分失败，1=全部成功
	ShareStatus int `json:"share_status,omitempty"`
	// SharePhotoExists 被推送视频是否已存在于被推送账户的视频库中，true = 已存在；false = 不存在
	SharePhotoExists bool `json:"share_photo_exists,omitempty"`
	// NotSupportedInternalPhoto 被推送账户是否是内部账户，true = 是，发；false = 不是
	NotSupportedInternalPhoto bool `json:"not_supported_internal_photo,omitempty"`
	// MismatchedAccountList 在“同主体同代理商”的基础上，本次新增“同主体账户”，“内部账户”等两种可推送账户类型，如果被推送账户不满足条件，将于此列表返回
	MismatchedAccountList []uint64 `json:"mismatched_account_list,omitempty"`
	// NeedToTryList 因各种原因导致推送失败，需要重试的推送记录
	NeedToTryList []PhotoShareResult `json:"need_to_try_list,omitempty"`
	// ShareSuccessList 视频推送成功结果网关参数
	ShareSuccessList []PhotoShareResult `json:"share_success_list,omitempty"`
}

// PhotoShareResult
type PhotoShareResult struct {
	// AccountID 账号ID
	AccountID uint64 `json:"account_id,omitempty"`
	// PhotoID 分享生成新的photoId
	PhotoID string `json:"photo_id,omitempty"`
	// OriginalPhotoID 被推送视频ID
	OriginalPhotoID string `json:"original_photo_id,omitempty"`
	// ShareResult 推送结果
	ShareResult string `json:"share_result,omitempty"`
}
