package file

import (
	"fmt"
	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/file"
	"testing"
)

func TestAdAppList(t *testing.T) {
	var (
		sdk = core.NewSDKClient(0, "xxxxx")
		req = &file.AdAppListRequest{
			AdvertiserID: 11044594,
		}
		list *file.AdAppListResponse
	)

	list, err := AdAppList(sdk, "36f7b5519e1f4437e382c6a34ed3aec3", req)
	fmt.Println(err)
	fmt.Println(list)
}
