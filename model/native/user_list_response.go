package native

type UserListResponse struct {
	UserList []*UserProfileViewSnake `json:"user_list"` //用户列表
}

type UserProfileViewSnake struct {
	UserID      int64  `json:"user_id"`       // 用户id
	UserName    string `json:"user_name"`     // 用户名称
	UserSex     string `json:"user_sex"`      // 用户性别，男性M，女性F
	HeadURL     string `json:"head_url"`      // 用户头像
	KOLUserType int    `json:"kol_user_type"` // 达人用户类型，2服务号达人，3聚星达人
}
