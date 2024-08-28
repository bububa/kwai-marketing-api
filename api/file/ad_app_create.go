package file

import (
	"context"

	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/file"
)

// AdAppCreate 创建应用
func AdAppCreate(ctx context.Context, clt *core.SDKClient, accessToken string, req *file.AdAppCreateRequest) (*file.App, error) {
	var resp file.App
	err := clt.Upload(ctx, accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
