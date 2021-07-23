package oauth

import "encoding/json"

// AccessTokenRequest 获取AccessToken APIRequest
type AccessTokenRequest struct {
	AppID    int64  `json:"app_id,omitempty"`
	Secret   string `json:"secret,omitempty"`
	AuthCode string `json:"auth_code,omitempty"`
}

// Url implement PostRequest interface
func (r AccessTokenRequest) Url() string {
	return "oauth2/authorize/access_token"
}

// Encode implement PostRequest interface
func (r AccessTokenRequest) Encode() []byte {
	buf, _ := json.Marshal(r)
	return buf
}
