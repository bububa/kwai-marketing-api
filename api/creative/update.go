package creative

import (
	"git.gametaptap.com/tapad/github/kwai-marketing-api/core"
	"git.gametaptap.com/tapad/github/kwai-marketing-api/model/creative"
)

// Update 修改创意
//【注】联盟广告不支持便利贴图片素材。
func Update(clt *core.SDKClient, accessToken string, req *creative.UpdateRequest) (int64, error) {
	var resp creative.UpdateResponse
	err := clt.Post(accessToken, req, &resp)
	if err != nil {
		return 0, err
	}
	return resp.CreativeID, nil
}
