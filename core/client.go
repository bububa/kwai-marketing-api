package core

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/bububa/kwai-marketing-api/core/internal/debug"
	"github.com/bububa/kwai-marketing-api/model"
)

var (
	onceInit   sync.Once
	httpClient *http.Client
)

func defaultHttpClient() *http.Client {
	onceInit.Do(func() {
		transport := http.DefaultTransport.(*http.Transport).Clone()
		transport.MaxIdleConns = 100
		transport.MaxConnsPerHost = 100
		transport.MaxIdleConnsPerHost = 100
		httpClient = &http.Client{
			Transport: transport,
			Timeout:   time.Second * 60,
		}
	})
	return httpClient
}

// SDKClient api client
type SDKClient struct {
	client *http.Client
	tracer *Otel
	secret string
	appID  uint64
	debug  bool
}

// NewSDKClient init sdk client
func NewSDKClient(appID uint64, secret string) *SDKClient {
	return &SDKClient{
		appID:  appID,
		secret: secret,
		client: defaultHttpClient(),
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

func (c *SDKClient) WithTracer(namespace string) {
	c.tracer = NewOtel(namespace, c.AppID())
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
func (c *SDKClient) Post(ctx context.Context, accessToken string, req model.PostRequest, resp interface{}) error {
	var reqResp model.Response
	if v, ok := resp.(model.Response); ok {
		reqResp = v
	} else {
		reqResp = &model.BaseResponse{}
	}
	err := c.post(ctx, accessToken, c.PostUrl(req), req.Encode(), reqResp)
	if err != nil {
		return err
	}
	if reqResp.IsError() {
		return reqResp
	}
	if resp != nil {
		if v, ok := reqResp.(*model.BaseResponse); ok {
			err = json.Unmarshal(v.Data, resp)
			if err != nil {
				return err
			}
		} else {
			resp = reqResp
		}
	}
	return nil
}

// Get execute get api request
func (c *SDKClient) Get(ctx context.Context, accessToken string, req model.GetRequest, resp interface{}) error {
	var reqResp model.BaseResponse
	err := c.get(ctx, accessToken, c.GetUrl(req), &reqResp)
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

// GetBytes get bytes api
func (c *SDKClient) GetBytes(ctx context.Context, accessToken string, req model.GetRequest) ([]byte, error) {
	reqUrl := c.GetUrl(req)
	debug.PrintGetRequest(reqUrl, c.debug)
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodGet, reqUrl, nil)
	if err != nil {
		return nil, err
	}
	if accessToken != "" {
		httpReq.Header.Add("Access-Token", accessToken)
	}
	var ret []byte
	c.WithSpan(ctx, httpReq, nil, nil, func(httpReq *http.Request, resp interface{}) (*http.Response, error) {
		httpResp, err := c.client.Do(httpReq)
		if err != nil {
			return httpResp, err
		}
		defer httpResp.Body.Close()
		ret, err = io.ReadAll(httpResp.Body)
		return httpResp, err
	})
	return ret, nil
}

func (c *SDKClient) GetOnBody(ctx context.Context, accessToken string, req model.PostRequest, resp interface{}) error {
	var reqResp model.BaseResponse
	err := c.getOnBody(ctx, accessToken, c.PostUrl(req), req.Encode(), &reqResp)
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
func (c *SDKClient) Upload(ctx context.Context, accessToken string, req model.UploadRequest, resp interface{}) error {
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
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, reqUrl, &buf)
	if err != nil {
		return err
	}
	httpReq.Header.Add("Content-Type", mw.FormDataContentType())
	if accessToken != "" {
		httpReq.Header.Add("Access-Token", accessToken)
	}

	var reqResp model.BaseResponse
	bs, _ := json.Marshal(mp)
	if err := c.WithSpan(ctx, httpReq, reqResp, bs, c.fetch); err != nil {
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
func (c *SDKClient) post(ctx context.Context, accessToken string, reqUrl string, reqBytes []byte, resp interface{}) error {
	debug.PrintPostJSONRequest(reqUrl, reqBytes, c.debug)
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, reqUrl, bytes.NewReader(reqBytes))
	if err != nil {
		return err
	}
	if accessToken != "" {
		httpReq.Header.Add("Access-Token", accessToken)
	}
	httpReq.Header.Add("Content-Type", "application/json")
	return c.WithSpan(ctx, httpReq, resp, reqBytes, c.fetch)
}

// get data through api
func (c *SDKClient) get(ctx context.Context, accessToken string, reqUrl string, resp interface{}) error {
	debug.PrintGetRequest(reqUrl, c.debug)
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodGet, reqUrl, nil)
	if err != nil {
		return err
	}
	if accessToken != "" {
		httpReq.Header.Add("Access-Token", accessToken)
	}
	httpReq.Header.Add("Content-Type", "application/json")
	return c.WithSpan(ctx, httpReq, resp, nil, c.fetch)
}

func (c *SDKClient) getOnBody(ctx context.Context, accessToken string, reqUrl string, reqBytes []byte, resp interface{}) error {
	debug.PrintPostJSONRequest(reqUrl, reqBytes, c.debug)
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodGet, reqUrl, bytes.NewReader(reqBytes))
	if err != nil {
		return err
	}
	if accessToken != "" {
		httpReq.Header.Add("Access-Token", accessToken)
	}
	httpReq.Header.Add("Content-Type", "application/json")
	return c.WithSpan(ctx, httpReq, resp, reqBytes, c.fetch)
}

func (c *SDKClient) fetch(httpReq *http.Request, resp interface{}) (*http.Response, error) {
	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return httpResp, err
	}
	defer httpResp.Body.Close()
	if err = debug.DecodeJSONHttpResponse(httpResp.Body, resp, c.debug); err != nil {
		debug.PrintError(err, c.debug)
		return httpResp, err
	}
	return httpResp, nil
}

func (c *SDKClient) WithSpan(ctx context.Context, req *http.Request, resp interface{}, payload []byte, fn func(*http.Request, interface{}) (*http.Response, error)) error {
	if c.tracer == nil {
		_, err := fn(req, resp)
		return err
	}
	return c.tracer.WithSpan(ctx, req, resp, payload, fn)
}
