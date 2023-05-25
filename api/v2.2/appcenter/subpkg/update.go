package subpkg

import (
	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/v2.2/appcenter/subpkg"
)

func Update(clt *core.SDKClient, accessToken string, req *subpkg.ListRequest) (*[]subpkg.SubpkgItem, error) {
	UpdateRequest := subpkg.UpdateRequest{
		AdvertiserID:    req.AdvertiserID,
		ParentPackageID: req.ParentPackageID,
		ChannelID:       req.ChannelID,
		Count:           req.Count,
		Type:            req.Type,
	}
	return nil, nil
}