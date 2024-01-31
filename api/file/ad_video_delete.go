package file

import (
	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/file"
)

// AdVideoDelete 批量删除视频
func AdVideoDelete(clt *core.SDKClient, accessToken string, req *file.AdVideoDeleteRequest) ([]string, error) {
	var ret file.AdVideoDeleteResponse
	if err := clt.Post(accessToken, req, &ret); err != nil {
		return nil, err
	}
	return ret.PhotoIDs, nil
}
