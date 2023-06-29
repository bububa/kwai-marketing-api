package native

import (
	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/native"
)

func UserList(clt *core.SDKClient, accessToken string, req *native.UserListRequest) (*native.UserListResponse, error) {
	var resp native.UserListResponse
	err := clt.Post(accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
