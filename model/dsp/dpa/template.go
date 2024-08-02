package dpa

// Template DPA 创意模板
type Template struct {
	// CoverURL 模板图片url。图片类模板表示图片地址；视频类模板表示视频对应的封面地址
	CoverURL string `json:"cover_url,omitempty"`
	// EditableAreaCoverURL 图片类模板的可编辑区域图片url
	EditableAreaCoverURL string `json:"editable_area_cover_url,omitempty"`
	// VideoURL 视频地址
	VideoURL string `json:"video_url,omitempty"`
	// Name 模板名称
	Name string `json:"name,omitempty"`
	// ID DPA 创意模板 id
	ID uint64 `json:"id,omitempty"`
	// VideoDuration 视频播放时长
	VideoDuration int64 `json:"video_duration,omitempty"`
	// Height 衍生物料的高
	Height int `json:"height,omitempty"`
	// Width 衍生物料的宽
	Width int `json:"width,omitempty"`
	// TemplateType 模板类型1:竖版图片、2:横版图片3:横版视频、4:竖版视频
	TemplateType int `json:"template_type,omitempty"`
}
