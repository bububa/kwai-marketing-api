package tool

// OperationRecord 账户操作记录
type OperationRecord struct {
	// ObjectID 操作对象ID
	ObjectID string `json:"object_id,omitempty"`
	// OperationType 操作类型; 1：新增 2：修改
	OperationType int `json:"operation_type,omitempty"`
	// OperationTarget 操作目标类型; 1：账户2：计划3：广告组4：创意5：视频6：应用7：人群包
	OperationTarget int `json:"operation_target,omitempty"`
	// RoleType 操作人; 1：广告主2：代理商3：系统4：管理员5：Market Api
	RoleType int `json:"role_type,omitempty"`
	// ObjectName 操作对象名称
	ObjectName string `json:"object_name,omitempty"`
	// OperationTime 操作时间
	OperationTime string `json:"operation_time,omitempty"`
	// ContentLog 日志内容
	ContentLog []ContentLog `json:"content_log,omitempty"`
}
