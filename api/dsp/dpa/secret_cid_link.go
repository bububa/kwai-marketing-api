package dpa

import (
	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/dsp/dpa"
)

// SecretCidLink CID服务商投放SDPA接口
func SecretCidLink(clt *core.SDKClient, accessToken string, req *dpa.SecretCidLinkRequest) error {
	return clt.Post(accessToken, req, nil)
}
