package native

import "encoding/json"

type PhotoListRequest struct {
	AuthorID     uint64 `json:"author_id"`     // 达人用户id
	AdvertiserID uint64 `json:"advertiser_id"` // 广告主ID
	PCursor      string `json:"pcursor"`       // 游标，第一次不传，后续滑动获取时根据结果返回的pcursor填取
	Count        int    `json:"count"`         // 每次获取的个数，最大不超过50个
	KOLUserType  int    `json:"kol_user_type"` // 原生达人类型，1普通快手号（备注：需要在“普通快手号原生白名单”中才能返回列表），2服务号原生，3聚星达人原生
	CampaignType int    `json:"campaign_type"` // 计划类型，原生场景下仅支持部分计划类型，2提升应用安装，5收集销售线索，7提升应用活跃，19小程序推广
	TabType      int    `json:"tab_type"`      // 0代表profile页非隐藏视频，1代表profile页隐藏视频
}

// Url implement PostRequest interface
func (r PhotoListRequest) Url() string {
	return "gw/dsp/v1/native/photo/list"
}

// Encode implement PostRequest interface
func (r PhotoListRequest) Encode() []byte {
	ret, _ := json.Marshal(r)
	return ret
}
