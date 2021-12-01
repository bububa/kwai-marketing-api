package oauth

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
	AdvertiserID int64 `json:"advertiser_id,omitempty"`
	// AdvertiserIDs 已授权账户所有的account_id
	AdvertiserIDs []int64 `json:"advertiser_ids,omitempty"`
}
