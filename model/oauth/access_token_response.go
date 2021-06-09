package oauth

type AccessTokenResponse struct {
	AccessToken           string  `json:"access_token,omitempty"`             // 用于验证权限的token
	AccessTokenExpiresIn  int64   `json:"access_token_expires_in,omitempty"`  // access_token剩余有效时间，单位：秒
	RefreshToken          string  `json:"refresh_token,omitempty"`            // 用于获取新的access_token和refresh_token，并且刷新过期时间
	RefreshTokenExpiresIn int64   `json:"refresh_token_expires_in,omitempty"` // refresh_token剩余有效时间，单位：秒
	AdvertiserID          int64   `json:"advertiser_id,omitempty"`            // 广告主ID
	AdvertiserIDs         []int64 `json:"advertiser_ids,omitempty"`           // 已授权账户所有的account_id
}
