package subpkg

import (
	"errors"

	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/appcenter/subpkg"
)

// Mod 【应用中心】更新/恢复/删除应用分包
func Mod(clt *core.SDKClient, accessToken string, req *subpkg.ModRequest) error {
	var resp subpkg.ModResponse
	if err := clt.Post(accessToken, req, &resp); err != nil {
		return err
	}
	if !resp.Result {
		return errors.New("更新/删除/恢复分包失败")
	}
	return nil
}
