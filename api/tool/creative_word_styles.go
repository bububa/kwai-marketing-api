package tool

import (
	"git.gametaptap.com/tapad/github/kwai-marketing-api/core"
	"git.gametaptap.com/tapad/github/kwai-marketing-api/model/tool"
)

// CreativeWordStyles 获取可选的封面贴纸样式
func CreativeWordStyles(clt *core.SDKClient, accessToken string, advertiserID int64) ([]tool.CreativeWordStyle, error) {
	req := &tool.CreativeWordStylesRequest{
		AdvertiserID: advertiserID,
	}
	var resp []tool.CreativeWordStyle
	err := clt.Get(accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
