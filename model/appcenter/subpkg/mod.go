package subpkg

import "github.com/bububa/kwai-marketing-api/model"

// ModRequest 【应用中心】更新/恢复/删除应用分包
type ModRequest struct {
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// PackageID 应用分包id
	PackageID []uint64 `json:"package_id,omitempty"`
	// PutStatus 操作分包的类型： 0-更新分包， 1-恢复分包， 2-删除分包
	// 当分包状态处于构建中或更新中时，不可对分包进行操作。更新分包指将分包升级到绑定最新发布版本的母包；删除分包将不能使用该分包创建广告；恢复分包可重新使用分包物料创建广告。
	PutStatus int `json:"put_status,omitempty"`
}

// Url implement PostRequest interface
func (r ModRequest) Url() string {
	return "gw/dsp/appcenter/subpkg/mod"
}

// Encode implement PostRequest interface
func (r ModRequest) Encode() []byte {
	return model.JSONMarshal(r)
}

// ModResponse 更新/恢复/删除应用分包响应 API Response
type ModResponse struct {
	// Result 更新/删除/恢复分包是否成功
	// true-成功，false-失败
	Result bool `json:"result"`
}
