package oauth

import "encoding/json"

type AccessTokenRequest struct {
	AppID    int64  `json:"app_id,omitempty"`
	Secret   string `json:"secret,omitempty"`
	AuthCode string `json:"auth_code,omitempty"`
}

func (r AccessTokenRequest) Url() string {
	return "oauth2/authorize/access_token"
}

func (r AccessTokenRequest) Encode() []byte {
	buf, _ := json.Marshal(r)
	return buf
}
