package video

import (
	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/dsp/video"
)

// QueryAutoShareSwitch 查询账户共享视频库按钮是否开启
func QueryAutoShareSwitch(clt *core.SDKClient, accessToken string, req *video.QueryAutoShareSwitchRequest) (*video.QueryAutoShareSwitchResponse, error) {
	var resp video.QueryAutoShareSwitchResponse
	if err := clt.Post(accessToken, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
