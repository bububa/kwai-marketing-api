package oauth

import "encoding/json"

// RefreshTokenRequest 刷新Token APIRequest
type RefreshTokenRequest struct {
	// AppID 应用ID
	AppID int64 `json:"app_id,omitempty"`
	// Secret 应用密钥
	Secret string `json:"secret,omitempty"`
	// RefreshToken 最近一次快手返回的refresh_token
	RefreshToken string `json:"refresh_token,omitempty"`
}

// Url implement PostRequest interface
func (r RefreshTokenRequest) Url() string {
	return "oauth2/authorize/refresh_token"
}

// Encode implement PostRequest inteface
func (r RefreshTokenRequest) Encode() []byte {
	buf, _ := json.Marshal(r)
	return buf
}
