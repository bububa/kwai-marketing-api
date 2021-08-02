package creative

import "encoding/json"

// BatchUpdateRequest 批量创建&修改创意 API Request
// 注：该接口可实现创意的批量创建&更新，上传creative_id为更新，不上传creative_id则为创建。
type BatchUpdateRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID int64 `json:"advertiser_id,omitempty"`
	// UnitID 广告组ID; 一个unit最多创建15个创意
	UnitID int64 `json:"unit_id,omitempty"`
	// ClickTrackUrl 第三方点击检测链接; 仅当广告组scene_id为1、2、6、7时，可选填； 广告组优化目标为激活时，该字段必填（下载类广告投放的应用集成快手Android SDK时除外） 使用Marketing API创建时，监测链接使用以该文档为准
	ClickTrackUrl string `json:"click_track_url,omitempty"`
	// ImpressionUrl 第三方开始播放监测链接; 仅当广告组scene_id为3时，可选填； 广告组优化目标为激活时，该字段必填（下载类广告投放的应用集成快手Android SDK时除外） 使用Marketing API创建时，监测链接使用以该文档为准
	ImpressionUrl string `json:"impression_url,omitempty"`
	// AdPhotoPlayedT3sUrl 第三方有效播放监测链接; 仅历史个别账户使用且当广告组scene_id为3时可选，与impression_url不可同时使用
	AdPhotoPlayedT3sUrl string `json:"ad_photo_played_t3s_url,omitempty"`
	// ActionbarClickUrl 第三方点击按钮监测链接; 1.校验click_url必填的广告场景 优选(1)/信息流(2、7)/上下滑（6） 2.优化目标为激活功能必填点击监测链接,但如果安卓应用接入了快手监测sdk就不需要填写监测链接了 3.联盟场景检查click_url不能为空 4.若广告联盟的转化目标为激活，click_url、actionbar_click_url和监测SDK至少三选一
	ActionbarClickUrl string `json:"actionbar_click_url,omitempty"`
	// CreativeCategory 创意分类; 由创意分类查询接口 获得；必须是叶子结点；与创意标签同时传或同时不传
	CreativeCategory int `json:"creative_category,omitempty"`
	// CreativeTag 创意标签; 与创意分类参数，要么都传，要么都不传；且单个创意的创意标签最多20个；单个创意标签不能为空且不能超过10字符
	CreativeTag []string `json:"creative_tag,omitempty"`
	// Creatives 批量参数
	Creatives []Creative `json:"creatives,omitempty"`
}

// Url implement PostRequest interface
func (r BatchUpdateRequest) Url() string {
	return "v2/creative/batch/update"
}

// Encode implement PostRequest interface
func (r BatchUpdateRequest) Encode() []byte {
	ret, _ := json.Marshal(r)
	return ret
}
