package oauth

import "encoding/json"

type RefreshTokenRequest struct {
	AppID        int64  `json:"app_id,omitempty"`
	Secret       string `json:"secret,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"` // 最近一次快手返回的refresh_token
}

func (r RefreshTokenRequest) Url() {
	return "oauth2/authorize/refresh_token"
}

func (r RefreshTokenRequest) Encode() []byte {
	buf, _ := json.Marshal(r)
	return buf
}
