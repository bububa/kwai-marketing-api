package file

import (
	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/file"
)

// AdVideoShareNew 视频库-推送视频(新版)
func AdVideoShareNew(clt *core.SDKClient, accessToken string, req *file.AdVideoShareNewRequest) (*file.AdVideoShareNewResponse, error) {
	var resp file.AdVideoShareNewResponse
	if err := clt.Post(accessToken, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
