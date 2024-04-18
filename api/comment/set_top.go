package comment

import (
	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/comment"
)

// SetTop 评论置顶
func SetTop(clt *core.SDKClient, accessToken string, req *comment.SetTopRequest) (uint64, error) {
	var resp comment.SetTopResponse
	if err := clt.Post(accessToken, req, &resp); err != nil {
		return 0, err
	}
	return resp.CommentID, nil
}
