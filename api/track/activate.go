package track

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/track"
)

// Activate 转化回传API
func Activate(req *track.ActivateRequest) error {
	reqUrl := fmt.Sprintf("%s?%s", core.ACTIVATE_URL, req.Encode())
	httpReq, err := http.NewRequest("GET", reqUrl, nil)
	httpResp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return err
	}
	defer httpResp.Body.Close()
	var ret struct {
		Result int `json:"result,omitempty"`
	}
	err = json.NewDecoder(httpResp.Body).Decode(&ret)
	if err != nil {
		return err
	}
	if ret.Result != 1 {
		return errors.New("activate failed")
	}
	return nil
}
