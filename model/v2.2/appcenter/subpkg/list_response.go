package subpkg

type ListResponse struct {
	CurrentPage int          `json:"current_page"`
	PageSize    int          `json:"page_size"`
	TotalCount  int          `json:"total_count"`
	List        []SubpkgItem `json:"list"`
}

type SubpkgItem struct {
	PackageID        int64  `json:"package_id"`         // 包ID
	ChannelID        string `json:"channel_id"`         // 渠道号(分包号)
	Description      string `json:"description"`        // 分包描述
	RealAppVersion   string `json:"real_app_version"`   // 应用版本信息
	SubPackageStatus int    `json:"sub_package_status"` // 应用分包状态: 1-审核中，2-审核失败，3-待发布，4-已发布，5-已下架 6-创建中，7-更新中，8-构建失败
	CanUpdate        bool   `json:"can_update"`         // 是否可更新: 仅分包管理列表时有效，表示应用分包是否可以更新
	UpdateTime       int64  `json:"update_time"`        // 更新时间: 仅分包管理列表时有效，表示应用分包的更新时间
	CanRecycle       *bool  `json:"can_recycle"`        // 是否可恢复: 仅分包回收站列表时有效，表示应用分包是否可以恢复
	DeleteTime       *int64 `json:"delete_time"`        // 删除时间: 仅分包回收站列表时有效，表示应用分包的删除时间
	ParentPackageID  int64  `json:"parent_package_id"`  // 分包的母包ID
}
