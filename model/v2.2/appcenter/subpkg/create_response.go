package subpkg

type CreateResponse struct {
	// CreativeID 创意ID
	Item []Item
}

type Item struct {
	PackageId       int64    `json:"package_id"`
	BuildStatus     int    `json:"build_status"`
	ParentPackageId int64  `json:"parent_package_id"`
	ChannelId       string `json:"channel_id"`
}

