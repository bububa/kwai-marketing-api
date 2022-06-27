package track

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/track"
)

// Activate 转化回传API
func Activate(req *track.ActivateRequest) error {
	reqUrl := fmt.Sprintf("%s?%s", core.ACTIVATE_URL, req.Encode())
	if req.Debug {
		log.Println("[DEBUG] [API] GET", reqUrl)
	}
	httpReq, err := http.NewRequest("GET", reqUrl, nil)
	if err != nil {
		return err
	}
	httpResp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return err
	}
	defer httpResp.Body.Close()
	var ret struct {
		Result int `json:"result,omitempty"`
	}
	if req.Debug {
		if body, err := ioutil.ReadAll(httpResp.Body); err != nil {
			return err
		} else {
			body2 := body
			buf := bytes.NewBuffer(make([]byte, 0, len(body2)+1024))
			if err := json.Indent(buf, body2, "", "    "); err == nil {
				body2 = buf.Bytes()
			}
			log.Printf("[DEBUG] [API] http response body:\n%s\n", body2)
			if err := json.Unmarshal(body, &ret); err != nil {
				return err
			}
		}
	} else if err := json.NewDecoder(httpResp.Body).Decode(&ret); err != nil {
		return err
	}
	if ret.Result != 1 {
		return errors.New("activate failed")
	}
	return nil
}
