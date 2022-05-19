package tool

import (
	"encoding/json"
	"net/url"
	"strconv"
)

// KeyFrameRequest 获取可选的推荐封面 API Request
type KeyFrameRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// PhotoIDs 视频ID; 1.目前只支持一次查询一个photo_id关键帧；2.视频的user_id要与当前account的user_id一致
	PhotoIDs []string `json:"photo_ids,omitempty"`
}

// Url implement GetRequest interface
func (r KeyFrameRequest) Url() string {
	return "v1/tool/key_frame"
}

// Encode implement GetRequest interface
func (r KeyFrameRequest) Encode() string {
	values := &url.Values{}
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	photoIDs, _ := json.Marshal(r.PhotoIDs)
	values.Set("photo_ids", string(photoIDs))
	return values.Encode()
}
