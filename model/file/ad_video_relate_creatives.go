package file

import "encoding/json"

// AdVideoRelateCreativesRequest 视频关联创意数查询
type AdVideoRelateCreativesRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// PhotoIDs 视频id; 最大20个，可以动态配置
	PhotoIds []string `json:"photo_ids,omitempty"`
}

// Url implement PostRequest interface
func (r AdVideoRelateCreativesRequest) Url() string {
	return "v1/file/ad/video/relate/creatives"
}

// Encode implement PostRequest interface3
func (r AdVideoRelateCreativesRequest) Encode() []byte {
	ret, _ := json.Marshal(r)
	return ret
}

// AdVideoRelateCreativesResponse 视频关联创意数查询 API Response
type AdVideoRelateCreativesResponse struct {
	// RelatedCreatives 与视频相关联的创意的信息; 关联创意数=审核中+投放中的创意，仅包含自定义创意，程序化创意数量暂未计入
	RelatedCreatives []AdVideoRelatedCreatives `json:"related_creatives,omitempty"`
}

// AdVideoRelatedCreatives 视频关联创意
type AdVideoRelatedCreatives struct {
	// PhotoID 视频id
	PhotoID string `json:"photo_id,omitempty"`
	// Creatives 与视频关联的创意的信息
	Creatives []AdVideoRelatedCreative `json:"creatives,omitempty"`
	// CreativeCount 与此视频id关联的创意总数
	CreativeCount int64 `json:"creative_count,omitempty"`
	// AdvancedCreativeIDs 与此视频 id 关联的程序化创意ID
	AdvancedCreativeIDs []uint64 `json:"advanced_creative_ids,omitempty"`
	// AdvancedCreativeCount 与此视频 id 关联的程序化创意数量
	AdvancedCreativeCount int64 `json:"advanced_creative_count,omitempty"`
}

// AdVideoRelatedCreative
type AdVideoRelatedCreative struct {
	// CreativeID 创意id
	CreativeID uint64 `json:"creative_id,omitempty"`
	// CreativeName 创意名称
	CreativeName string `json:"creative_name,omitempty"`
}
