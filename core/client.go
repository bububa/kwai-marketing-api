package core

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"strings"

	"github.com/bububa/kwai-marketing-api/core/internal/debug"
	"github.com/bububa/kwai-marketing-api/model"
)

// SDKClient api client
type SDKClient struct {
	appID  uint64
	secret string
	debug  bool
	client *http.Client
}

// NewSDKClient init sdk client
func NewSDKClient(appID uint64, secret string) *SDKClient {
	return &SDKClient{
		appID:  appID,
		secret: secret,
		client: http.DefaultClient,
	}
}

// SetDebug set debug mode
func (c *SDKClient) SetDebug(debug bool) {
	c.debug = debug
}

// AppID appID getter
func (c *SDKClient) AppID() uint64 {
	return c.appID
}

// Secret secret getter
func (c *SDKClient) Secret() string {
	return c.secret
}

func (c *SDKClient) SetHttpClient(client *http.Client) {
	c.client = client
}

// PostUrl post请求地址
func (c *SDKClient) PostUrl(req model.PostRequest) string {
	return fmt.Sprintf("%s/%s", BASE_URL, req.Url())
}

// GetUrl get请求地址
func (c *SDKClient) GetUrl(req model.GetRequest) string {
	return fmt.Sprintf("%s/%s?%s", BASE_URL, req.Url(), req.Encode())
}

// UploadUrl post multipart/form-data请求地址
func (c *SDKClient) UploadUrl(req model.UploadRequest) string {
	return fmt.Sprintf("%s/%s", BASE_URL, req.Url())
}

// Post execute post api request
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

// Get execute get api request
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

func (c *SDKClient) GetOnBody(accessToken string, req model.PostRequest, resp interface{}) error {
	var reqResp model.BaseResponse
	err := c.getOnBody(accessToken, c.PostUrl(req), req.Encode(), &reqResp)
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

// Upload multipart/form-data post
func (c *SDKClient) Upload(accessToken string, req model.UploadRequest, resp interface{}) error {
	var buf bytes.Buffer
	mw := multipart.NewWriter(&buf)
	params := req.Encode()
	mp := make(map[string]string, len(params))
	for _, v := range params {
		var (
			fw  io.Writer
			r   io.Reader
			err error
		)
		if v.Reader != nil {
			if fw, err = mw.CreateFormFile(v.Key, v.Value); err != nil {
				return err
			}
			r = v.Reader
			mp[v.Key] = fmt.Sprintf("@%s", v.Value)
		} else {
			if fw, err = mw.CreateFormField(v.Key); err != nil {
				return err
			}
			r = strings.NewReader(v.Value)
			mp[v.Key] = v.Value
		}
		if _, err = io.Copy(fw, r); err != nil {
			return err
		}
	}
	mw.Close()
	reqUrl := c.UploadUrl(req)
	debug.PrintPostMultipartRequest(reqUrl, mp, c.debug)
	httpReq, err := http.NewRequest("POST", reqUrl, &buf)
	if err != nil {
		return err
	}
	httpReq.Header.Add("Content-Type", mw.FormDataContentType())
	if accessToken != "" {
		httpReq.Header.Add("Access-Token", accessToken)
	}

	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return err
	}
	defer httpResp.Body.Close()
	var reqResp model.BaseResponse
	err = debug.DecodeJSONHttpResponse(httpResp.Body, &reqResp, c.debug)
	if err != nil {
		debug.PrintError(err, c.debug)
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
	httpResp, err := c.client.Do(httpReq)
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
	httpResp, err := c.client.Do(httpReq)
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

func (c *SDKClient) getOnBody(accessToken string, reqUrl string, reqBytes []byte, resp interface{}) error {
	debug.PrintGetRequest(reqUrl, c.debug)
	httpReq, err := http.NewRequest("GET", reqUrl, bytes.NewReader(reqBytes))
	if err != nil {
		return err
	}
	if accessToken != "" {
		httpReq.Header.Add("Access-Token", accessToken)
	}
	httpReq.Header.Add("Content-Type", "application/json")
	httpResp, err := c.client.Do(httpReq)
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
