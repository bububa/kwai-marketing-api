package region

// ListRequest 获取地域定向
type ListRequest struct{}

// Url implement GetRequest interface
func (r ListRequest) Url() string {
	return "v1/region/list"
}

// Encode implement GetRequest interface
func (r ListRequest) Encode() string {
	return ""
}
