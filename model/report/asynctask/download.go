package asynctask

import (
	"net/url"
	"strconv"
)

// DownloadRequest 下载任务结果
type DownloadRequest struct {
	// AdvertiserID 广告主 ID（注：非账户快手 ID），在获取 accessToken 时返回
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// TaskID 任务 ID
	TaskID uint64 `json:"task_id,omitempty"`
}

// Url implement GetRequest interface
func (r DownloadRequest) Url() string {
	return "v1/async_task/download"
}

// Encode implement GetRequest interface
func (r DownloadRequest) Encode() string {
	values := &url.Values{}
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	values.Set("task_id", strconv.FormatUint(r.TaskID, 10))
	return values.Encode()
}
