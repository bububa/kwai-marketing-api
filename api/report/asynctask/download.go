package asynctask

import (
	"context"

	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/report/asynctask"
)

func Download(ctx context.Context, clt *core.SDKClient, accessToken string, req *asynctask.DownloadRequest) ([]byte, error) {
	return clt.GetBytes(ctx, accessToken, req)
}
