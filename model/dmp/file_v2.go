package dmp

type FileV2 struct {
	// FilePath 文件路径 	包含作为文件唯一标识的字符串
	FilePath string `json:"file_path,omitempty"`
	// UploadFileType 上传文件格式 	TXT或ZIP
	UploadFileType string `json:"upload_file_type,omitempty"`
	// RecordSize 上传文件总行数
	RecordSize int64 `json:"record_size,omitempty"`
	// MatchType 文件数据类型（英文）
	MatchType string `json:"match_type,omitempty"`
	// Type 文件数据类型  1：IMEI；2：IDFA；3：IMEI_MD5；4：IDFA_MD5；5：手机号-MD5；7：OAID；8：OAID_MD5 9: 手机号_SHA256
	Type int64 `json:"type,omitempty"`
	// FileSize 文件大小
	FileSize int64 `json:"file_size,omitempty"`
	// Md5 文件摘要标识
	Md5 string `json:"md5,omitempty"`
	// AdvertiserId 广告主ID
	AdvertiserId int64 `json:"advertiser_id,omitempty"`
}
