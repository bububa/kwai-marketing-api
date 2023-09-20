package native

type PhotoListResponse struct {
	Photos  []*KwaiOrdePhotoViewSnake `json:"photos"`  // 视频列表
	PCursor string                    `json:"pcursor"` // 下标，如果后面无更多视频则返回 no_more; 如果后续有更多视频需用此返回填到拉取视频接口的入参上
}

type CdnUrlInfo struct {
	Cdn string `json:"cdn"` // cdn信息
	Url string `json:"url"` // url信息
}

type KwaiOrdePhotoViewSnake struct {
	PhotoID              string       `json:"photo_id"`                     // 加密后的photoId
	Caption              string       `json:"caption"`                      // 视频标题
	CoverURL             []CdnUrlInfo `json:"cover_url"`                    // 封面url
	MovieURL             []CdnUrlInfo `json:"movie_url"`                    // 视频url
	Duration             int64        `json:"duration"`                     // 视频时长，单位毫秒
	Height               int          `json:"height"`                       // 视频高度
	Width                int          `json:"width"`                        // 视频宽度
	CreativeMaterialType int          `json:"creative_material_type"`       // 视频横竖版: 1竖版，2横版
	AdSocialOrderID      int          `json:"ad_socail_order_id,omitempty"` // 聚星订单id
}
