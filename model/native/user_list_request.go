package native

import "encoding/json"

type UserListRequest struct {
	AdvertiserID uint64 `json:"advertiser_id"` // 广告主ID
	KOLUserType  []int  `json:"kol_user_type"` // 达人原生类型，1代表普通快手号（备注：需要在“普通快手号原生白名单”中才能返回列表），2服务号原生达人，3聚星原生达人
}

// Url implement PostRequest interface
func (r UserListRequest) Url() string {
	return "gw/dsp/v1/native/user/list"
}

// Encode implement PostRequest interface
func (r UserListRequest) Encode() []byte {
	ret, _ := json.Marshal(r)
	return ret
}
