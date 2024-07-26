package comment

import (
	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/comment"
)

// Reply 回复评论
func Reply(clt *core.SDKClient, accessToken string, req *comment.ReplyRequest) ([]comment.ReplyResult, error) {
	var resp comment.ReplyResponse
	err := clt.Post(accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return resp.ReplyResultList, nil
}
