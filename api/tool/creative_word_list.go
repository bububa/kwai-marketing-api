package tool

import (
	"github.com/Shinku-Chen/kwai-marketing-api/core"
	"github.com/Shinku-Chen/kwai-marketing-api/model/tool"
)

// CreativeWordList 获取可选的动态词包
func CreativeWordList(clt *core.SDKClient, accessToken string, advertiserID int64) ([]tool.CreativeWord, error) {
	req := &tool.CreativeWordListRequest{
		AdvertiserID: advertiserID,
	}
	var resp []tool.CreativeWord
	err := clt.Get(accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
