package jingbell

import (
	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model"
	"github.com/bububa/kwai-marketing-api/model/dsp/jingbell"
)

// Share 小铃铛推送
func Share(clt *core.SDKClient, accessToken string, req *jingbell.ShareRequest) error {
	var resp jingbell.ShareResponse
	if err := clt.Post(accessToken, req, &resp); err != nil {
		return err
	}
	if resp.Result != 1 {
		return model.BaseResponse{Code: resp.Result, Message: resp.Data}
	}
	return nil
}
