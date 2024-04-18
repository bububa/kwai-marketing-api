package comment

import (
	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/comment"
)

// ShieldInfoDelete 删除屏蔽评论信息
func ShieldInfoDelete(clt *core.SDKClient, accessToken string, req *comment.ShieldInfoDeleteRequest) error {
	return clt.Post(accessToken, req, nil)
}
