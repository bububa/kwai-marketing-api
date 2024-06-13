package asynctask

import "github.com/bububa/kwai-marketing-api/model"

// TaskParams 任务参数
type TaskParams struct {
	// StartDate 查询开始日期 格式如：yyyy-MM-dd，时间跨度不能超过 6 个月
	StartDate string `json:"start_date,omitempty"`
	// EndDate 查询结束日期 格式如：yyyy-MM-dd，不大于查询开始日期
	EndDate string `json:"end_date,omitempty"`
	// 查询维度 1：账户维度查询，2: 广告计划维度查询，3：广告组维度查询，4：广告创意维度查询 5：视频报表 报表 7：封面报表 8：便利贴报表 10：程序化创意 2.0 报表
	ViewType int `json:"view_type,omitempty"`
	// CampaignIDs 广告计划 ID 集
	CampaignIDs []uint64 `json:"campaign_ids,omitempty"`
	// UnitIDs 广告组ID集
	UnitIDs []uint64 `json:"unit_ids,omitempty"`
	// CreativeIDs 广告创意 ID 集
	CreativeIDs []uint64 `json:"creative_ids,omitempty"`
	// WordInfoIDs 推广关键词ID集，过滤筛选条件，单次查询数量不超过 5000,推广关键词ID集可通过获取关键词列表接口获取
	WordInfoIDs []uint64 `json:"word_info_ids,omitempty"`
	// Query 用户搜索词，过滤筛选条件，单次查询数量不超过5000
	Query []string `json:"query,omitempty"`
	// PhotoIDs 视频ID集 仅 view_type=5、7、8 可使用
	PhotoIDs []string `json:"photo_ids,omitempty"`
	// CoverIDs 封面ID集 仅 view_type=5、7、8 可使用
	CoverIDs []uint64 `json:"cover_ids,omitempty"`
	// VirtualCreativeIDs 程序化创意 ID集 进行 virtual_creative_id 的筛选
	VirtualCreativeIDs []uint64 `json:"virtual_creative_ids,omitempty"`
	// ReportDims "adScene"：按广告场景；"placementType"：按广告范围，快手/联盟;不传/传空/传空数组：不限
	ReportDims []string `json:"report_dims,omitempty"`
	// TemporalGranularity 时间粒度 “DAILY”：天粒度；“HOURLY”：小时粒度；默认按天粒度进行聚合
	TemporalGranularity model.TemporalGranularityType `json:"temporal_granularity,omitempty"`
	// SelectedColumns 自定义列	仅view_type=21或25支持使用，支持列名及其关联字段详见：https://docs.qingque.cn/d/home/eZQB-fLBIZLvGG50L7vFkHL3J?identityId=1oE314hFZmG
	SelectedColumns []string `json:"selected_columns,omitempty"`
}

type Task struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// TaskID 任务ID
	TaskID uint64 `json:"task_id,omitempty"`
	// TaskName 任务名称
	TaskName string `json:"task_name,omitempty"`
	// CreateTime 任务创建时间 格式如：yyyy-MM-dd HH:mm:ss
	CreateTime string `json:"create_time,omitempty"`
	// TaskStatus 任务状态 0：新建，1：处理中，2：处理成功，3：处理失败
	TaskStatus int `json:"task_status,omitempty"`
	// FileSize 文件大小 字节数
	FileSize int64 `json:"file_size,omitempty"`
}
