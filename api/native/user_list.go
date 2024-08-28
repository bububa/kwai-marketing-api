package native

import (
	"context"

	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/native"
)

func UserList(ctx context.Context, clt *core.SDKClient, accessToken string, req *native.UserListRequest) (*native.UserListResponse, error) {
	var resp native.UserListResponse
	err := clt.Post(ctx, accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
