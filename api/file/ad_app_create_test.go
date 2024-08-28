package file

import (
	"context"
	"fmt"
	"testing"

	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/file"
)

func CreatKuaishouApp(ctx context.Context) (app *file.App, err error) {
	sdk := core.NewSDKClient(0, "xxxxx")
	req := &file.AdAppCreateRequest{
		AdvertiserID:        11044594,
		AppName:             "测试11044594-83022",
		URL:                 "wwww.baidu.com",
		PackageName:         "com.taptap",
		AppVersion:          "2.22.0-mkt.30000123-14",
		AppPrivacyUrl:       "https://wwww.baidu.com",
		Platform:            1,
		ImageToken:          "c11a7c41163041e4b1f3e3fbf48f8a91.png",
		AppDetailImageToken: "38146cb2bb334d6f99f30f1ca096b86c.jpg",
	}
	app, err = AdAppCreate(ctx, sdk, "36f7b5519e1f4437e382c6a34ed3aec3", req)
	fmt.Println(err)
	return
}

func TestData(t *testing.T) {
	app, _ := CreatKuaishouApp(context.Background())
	fmt.Println(app)
}
