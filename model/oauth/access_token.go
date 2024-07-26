package oauth

import "encoding/json"

// AccessTokenRequest 获取AccessToken APIRequest
type AccessTokenRequest struct {
	// AppID 申请应用后快手返回的 app_id
	AppID uint64 `json:"app_id,omitempty"`
	// Secret 申请应用后快手返回的 secret
	Secret string `json:"secret,omitempty"`
	// AuthCode 授权时返回的 auth_code
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

// AccessTokenResponse 获取AccessToken APIResponse
type AccessTokenResponse struct {
	// AccessToken 用于验证权限的token
	AccessToken string `json:"access_token,omitempty"`
	// AccessTokenExpiresIn access_token剩余有效时间，单位：秒
	AccessTokenExpiresIn int64 `json:"access_token_expires_in,omitempty"`
	// RefreshToken 用于获取新的access_token和refresh_token，并且刷新过期时间
	RefreshToken string `json:"refresh_token,omitempty"`
	// RefreshTokenExpiresIn refresh_token剩余有效时间，单位：秒
	RefreshTokenExpiresIn int64 `json:"refresh_token_expires_in,omitempty"`
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// AdvertiserIDs 已授权账户所有的account_id
	AdvertiserIDs []uint64 `json:"advertiser_ids,omitempty"`
}
