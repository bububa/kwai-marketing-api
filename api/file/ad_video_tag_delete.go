package file

import (
	"github.com/Shinku-Chen/kwai-marketing-api/core"
	"github.com/Shinku-Chen/kwai-marketing-api/model/file"
)

// AdVideoTagDelete 视频库-删除视频标签
func AdVideoTagDelete(clt *core.SDKClient, accessToken string, req *file.AdVideoTagDeleteRequest) error {
	return clt.Post(accessToken, req, nil)
}
