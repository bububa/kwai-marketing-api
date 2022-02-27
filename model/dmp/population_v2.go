package dmp

// Population 人群包
type PopulationV2 struct {
	Population `json:",inline"`
	// FailedFilePaths 创建失败的文件路径列表
	FailedFilePaths []string `json:"failed_file_paths"`
	AdvertiserId    int64    `json:"advertiser_id"`
}
