package target

import "github.com/bububa/kwai-marketing-api/model"

// OptionDistanceListRequest 根据店铺名称查询商圈信息 API Request
type OptionDistanceListRequest struct {
	// ProvinceName 省份，城市二选一必填
	ProvinceName string `json:"province_name,omitempty"`
	// CityName 省份，城市二选一必填
	CityName string `json:"city_name,omitempty"`
	// DistrictName 街区，选填
	DistrictName string `json:"district_name,omitempty"`
	// LocationName 店铺名，必填
	LocationName string `json:"location_name,omitempty"`
	// Page 页数
	Page int `json:"page,omitempty"`
	// PageSize 页面大小
	PageSize int `json:"page_size,omitempty"`
}

// Url implement GetRequest interface
func (r OptionDistanceListRequest) Url() string {
	return "gw/dsp/target/option/distance_list"
}

// Encode implement PostRequest interface
func (r OptionDistanceListRequest) Encode() []byte {
	return model.JSONMarshal(r)
}

// OptionDistanceListResponse 根据店铺名称查询商圈信息 API Response
type OptionDistanceListResponse struct {
	// List 商圈信息列表
	List []Distance `json:"list,omitempty"`
	// CurrentPage 当前页数
	CurrentPage int `json:"current_page,omitempty"`
	// PageSize 页面大小
	PageSize int `json:"page_size,omitempty"`
	// NextPage 下一页页数
	NextPage int `json:"next_page,omitempty"`
}

// Distance 商圈信息
type Distance struct {
	// Address 地址
	Address string `json:"address,omitempty"`
	// Lat 经度
	Lat string `json:"lat,omitempty"`
	// Lng 纬度
	Lng string `json:"lng,omitempty"`
	// LocationName 店铺名称
	LocationName string `json:"location_name,omitempty"`
	// PoiID 唯一标识
	PoiID string `json:"poi_id,omitempty"`
	// Radius 半径（单位m）
	Radius int64 `json:"radius,omitempty"`
}
