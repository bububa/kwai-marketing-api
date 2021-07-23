package oauth

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

// UrlRequest 生成授权链接APIRequest
type UrlRequest struct {
	// AppID 应用ID
	AppID int64 `json:"app_id,omitempty"`
	// Scope 授权scope
	Scope []string `json:"scope,omitempty"`
	// RedirectUri 申请应用时开发者提供的回调地址，使用时需要UrlEncode一次
	RedirectUri string `json:"redirect_uri,omitempty"`
	// State 自定义参数; 回调时会原样返回，可用于广告主区分不同投放渠道等用途，广告主可选择性使用
	State string `json:"state,omitempty"`
	// OauthType 代理商使用授权URL拼接&oauth_type=agent
	OauthType string `json:"oauth_type,omitempty"`
}

// Encode implement GetRequest interface
func (r UrlRequest) Encode() string {
	values := url.Values{}
	values.Set("app_id", strconv.FormatInt(r.AppID, 10))
	values.Set("redirect_uri", r.RedirectUri)
	if len(r.Scope) > 0 {
		var scopes []string
		for _, s := range r.Scope {
			scopes = append(scopes, fmt.Sprintf(`"%s"`, s))
		}
		values.Set("scope", strings.Join(scopes, ","))
	}
	if r.State != "" {
		values.Set("state", r.State)
	}
	if r.OauthType != "" {
		values.Set("oauth_type", r.OauthType)
	}
	return values.Encode()
}

// Url implement GetRequest interface
func (r UrlRequest) Url() string {
	return "oauth"
}
