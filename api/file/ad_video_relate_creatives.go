package file

import (
	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/file"
)

// AdVideoRelateCreatives 视频关联创意数查询
func AdVideoRelateCreatives(clt *core.SDKClient, accessToken string, req *file.AdVideoRelateCreativesRequest) ([]file.AdVideoRelatedCreatives, error) {
	var resp file.AdVideoRelateCreativesResponse
	err := clt.Post(accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return resp.RelatedCreatives, nil
}
