package core

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/bububa/kwai-marketing-api/core/internal/debug"
	"github.com/bububa/kwai-marketing-api/model"
)

// sdkclient object
type SDKClient struct {
	appID  int64
	secret string
	debug  bool
}

// init sdk client
func NewSDKClient(appID int64, secret string) *SDKClient {
	return &SDKClient{
		appID:  appID,
		secret: secret,
	}
}

// set debug mode
func (c *SDKClient) SetDebug(debug bool) {
	c.debug = debug
}

// appID getter
func (c *SDKClient) AppID() int64 {
	return c.appID
}

// secret getter
func (c *SDKClient) Secret() string {
	return c.secret
}

func (c *SDKClient) PostUrl(req model.PostRequest) string {
	return fmt.Sprintf("%s/%s", BASE_URL, req.Url())
}

func (c *SDKClient) GetUrl(req model.GetRequest) string {
	return fmt.Sprintf("%s/%s?%s", BASE_URL, req.Url(), req.Encode())
}

// execute post api request
func (c *SDKClient) Post(accessToken string, req model.PostRequest, resp interface{}) error {
	var reqResp model.BaseResponse
	err := c.post(accessToken, c.PostUrl(req), req.Encode(), &reqResp)
	if err != nil {
		return err
	}
	if reqResp.IsError() {
		return reqResp
	}
	if resp != nil {
		err = json.Unmarshal(reqResp.Data, resp)
		if err != nil {
			return err
		}
	}
	return nil
}

// execute get api request
func (c *SDKClient) Get(accessToken string, req model.GetRequest, resp interface{}) error {
	var reqResp model.BaseResponse
	err := c.get(accessToken, c.GetUrl(req), &reqResp)
	if err != nil {
		return err
	}
	if reqResp.IsError() {
		return reqResp
	}
	if resp != nil {
		err = json.Unmarshal(reqResp.Data, resp)
		if err != nil {
			return err
		}
	}
	return nil
}

// post data through api
func (c *SDKClient) post(accessToken string, reqUrl string, reqBytes []byte, resp interface{}) error {
	debug.PrintPostJSONRequest(reqUrl, reqBytes, c.debug)
	httpReq, err := http.NewRequest("POST", reqUrl, bytes.NewReader(reqBytes))
	if err != nil {
		return err
	}
	if accessToken != "" {
		httpReq.Header.Add("Access-Token", accessToken)
	}
	httpReq.Header.Add("Content-Type", "application/json")
	httpResp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return err
	}
	defer httpResp.Body.Close()
	err = debug.DecodeJSONHttpResponse(httpResp.Body, resp, c.debug)
	if err != nil {
		debug.PrintError(err, c.debug)
		return err
	}
	return nil
}

// get data through api
func (c *SDKClient) get(accessToken string, reqUrl string, resp interface{}) error {
	debug.PrintGetRequest(reqUrl, c.debug)
	httpReq, err := http.NewRequest("GET", reqUrl, nil)
	if err != nil {
		return err
	}
	if accessToken != "" {
		httpReq.Header.Add("Access-Token", accessToken)
	}
	httpReq.Header.Add("Content-Type", "application/json")
	httpResp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return err
	}
	defer httpResp.Body.Close()
	err = debug.DecodeJSONHttpResponse(httpResp.Body, resp, c.debug)
	if err != nil {
		debug.PrintError(err, c.debug)
		return err
	}
	return nil
}
