package native

import (
	"context"

	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/dsp/native"
)

// OpenAccountNative 开启原生扩量开关接口
func OpenAccountNative(ctx context.Context, clt *core.SDKClient, accessToken string, req *native.OpenAccountNativeRequest) error {
	var resp native.OpenAccountNativeResponse
	return clt.Post(ctx, accessToken, req, &resp)
}
