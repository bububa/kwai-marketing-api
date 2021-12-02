package file

import (
	"git.gametaptap.com/tapad/github/kwai-marketing-api/core"
	"git.gametaptap.com/tapad/github/kwai-marketing-api/model/file"
)

// AdVideoShare 视频库-推送视频
func AdVideoShare(clt *core.SDKClient, accessToken string, req *file.AdVideoShareRequest) ([]file.AdVideoShareDetail, error) {
	var resp file.AdVideoShareResponse
	err := clt.Post(accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return resp.Details, nil
}
