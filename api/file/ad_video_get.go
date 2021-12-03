package file

import (
	"github.com/Shinku-Chen/kwai-marketing-api/core"
	"github.com/Shinku-Chen/kwai-marketing-api/model/file"
)

// AdVideoGet 查询视频信息get接口
func AdVideoGet(clt *core.SDKClient, accessToken string, req *file.AdVideoGetRequest) ([]file.Video, error) {
	var resp []file.Video
	err := clt.Post(accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
