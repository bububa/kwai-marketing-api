package unit

// DpaUnitParam DPA相关商品信息
type DpaUnitParam struct {
	// LibraryID 商品库ID; 当子计划类型为sdpa时必填
	LibraryID uint64 `json:"library_id,omitempty"`
	// OuterID 外部商品ID
	OuterID string `json:"outer_id,omitempty"`
	// DetailUnitType 商品链接类型; 0：app下载，1：H5跳转，2：Deeplink唤起，当计划类型为sdpa时必填
	DetailUnitType int `json:"detail_unit_type"`
	// ProductID 快手商品ID; 当计划子类型为sdpa时必填
	ProductID string `json:"product_id,omitempty"`
	// ClickUrl 点击跳转监控链接
	ClickUrl string `json:"click_url,omitempty"`
	// ActionbarClickUrl actionbar点击跳转监控链接
	ActionbarClickUrl string `json:"actionbar_click_url,omitempty"`
	// ImpressionUrl 曝光监控链接;
	ImpressionUrl string `json:"impression_url,omitempty"`
}
