package creative

import (
	"github.com/Shinku-Chen/kwai-marketing-api/core"
	"github.com/Shinku-Chen/kwai-marketing-api/model/creative"
)

// Preview 创意体验
func Preview(clt *core.SDKClient, accessToken string, req *creative.PreviewRequest) error {
	return clt.Post(accessToken, req, nil)
}
