package file

import (
	"context"
	"fmt"
	"testing"

	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/file"
)

func TestAdAppList(t *testing.T) {
	var (
		sdk = core.NewSDKClient(0, "xxxxx")
		req = &file.AdAppListRequest{
			AdvertiserID: 11044594,
		}
		list *file.AdAppListResponse
	)

	list, err := AdAppList(context.Background(), sdk, "36f7b5519e1f4437e382c6a34ed3aec3", req)
	fmt.Println(err)
	fmt.Println(list)
}
