package unit

type BackflowForecast struct {
	// CvLower 回流预估值的下限
	CvLower int64 `json:"backflow_cv_lower,omitempty"`
	// CvUpper 回流预估值的上限
	CvUpper int64 `json:"backflow_cv_upper,omitempty"`
	// Timestamp 本次回流预估数据的时间戳，13位毫秒时间戳
	Timestamp int64 `json:"backflow_timestamp,omitempty"`
}
