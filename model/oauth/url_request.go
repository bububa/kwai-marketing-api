package oauth

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

type UrlRequest struct {
	AppID       int64    `json:"app_id,omitempty"` // 应用ID
	Scope       []string `json:"scope,omitempty"`
	RedirectUri string   `json:"redirect_uri,omitempty"` // 申请应用时开发者提供的回调地址，使用时需要UrlEncode一次
	State       string   `json:"state,omitempty"`        // 自定义参数; 回调时会原样返回，可用于广告主区分不同投放渠道等用途，广告主可选择性使用
	OauthType   string   `json:"oauth_type,omitempty"`   // 代理商使用授权URL拼接&oauth_type=agent
}

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

func (r UrlRequest) Url() string {
	return "oauth"
}
