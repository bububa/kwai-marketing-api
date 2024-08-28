package file

import (
	"context"

	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/file"
)

// AdAppUpdate 修改应用
func AdAppUpdate(ctx context.Context, clt *core.SDKClient, accessToken string, req *file.AdAppUpdateRequest) (*file.App, error) {
	var resp file.App
	err := clt.Upload(ctx, accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
