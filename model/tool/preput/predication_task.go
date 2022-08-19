package preput

// AdPredicationTask 投前预估创建任务
type AdPredicationTask struct {
	// PhotoID 视频id
	PhotoID string `json:"photo_id,omitempty"`
	// Filename 视频文件名
	Filename string `json:"file_name,omitempty"`
	// OcpxActionType 优化目标，53表示表单提交，180表示激活
	OcpxActionType int `json:"ocpx_action_type,omitempty"`
}
