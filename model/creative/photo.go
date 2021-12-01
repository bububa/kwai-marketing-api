package creative

// Photo 新创建程序化创意请使用素材列表
type Photo struct {
	// PhotoID 视频ID
	PhotoID int64 `json:"photo_id,omitempty"`
	// CoverImageToken 封面图片token;通过上传图片接口获得，不传值则直接使用视频的首帧作为封面图片，自定义封面的图片宽高要与视频宽高一致，使用智能抽帧时不需要传递。
	CoverImageToken string `json:"cover_image_token,omitempty"`
	// CreativeMaterialType 素材类型; 1：竖版视频 2：横版视频
	CreativeMaterialType int `json:"creative_material_type,omitempty"`
}
