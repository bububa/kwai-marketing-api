package app

// Create 创建创意
// 【注】联盟广告不支持便利贴图片素材，只有联盟广告支持横版竖版图片素材。
func List(clt *core.SDKClient, accessToken string, req *app.ListRequest) ([]app.Item, error) {
	var resp app.ListResponse
	err := clt.Post(accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

