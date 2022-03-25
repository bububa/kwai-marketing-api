package file

// AdVideoShareResponse 视频库-推送视频 API Response
type AdVideoShareResponse struct {
	// Details Response details
	Details []AdVideoShareDetail `json:"details,omitempty"`
}

// AdVideoShareDetail 视频库-推送视频 API Response Detail
type AdVideoShareDetail struct {
	// AdvertiserID 账号ID
	AdvertiserID int64 `json:"advertiser_id,omitempty"`
	// PhotoID 分享生成新的photoId
	PhotoID string `json:"photo_id,omitempty"`
	// OriginalPhotoID 原始photoId
	OriginalPhotoID string `json:"original_photo_id,omitempty"`
}
