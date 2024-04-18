package comment

import (
	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/comment"
)

// ShieldInfoCreate 增加屏蔽评论信息
func ShieldInfoCreate(clt *core.SDKClient, accessToken string, req *comment.ShieldInfoCreateRequest) ([]uint64, error) {
	var resp comment.ShieldInfoCreateResponse
	if err := clt.Post(accessToken, req, &resp); err != nil {
		return nil, err
	}
	return resp.ShieldInfoIDList, nil
}
