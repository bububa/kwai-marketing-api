package unit

// Schedule 投放时段 不投放的时段为null，详请见下方，历史字段，即将废弃
type Schedule struct {
	// Mon 周一时间段; 时间段范围0, 23
	Mon []int `json:"mon,omitempty"`
	// Tues 周二时间段; 时间段范围0, 23
	Tues []int `json:"tues,omitempty"`
	// Wed 周三时间段; 时间段范围0, 23
	Wed []int `json:"wed,omitempty"`
	// Thur 周四时间段; 时间段范围0, 23
	Thur []int `json:"thur,omitempty"`
	// Fri 周五时间段; 时间段范围0, 23
	Fir []int `json:"fir,omitempty"`
	// Sat 周六时间段; 时间段范围0, 23
	Sat []int `json:"sat,omitempty"`
	// Sun 周日时间段; 时间段范围0, 23
	Sun []int `json:"sun,omitempty"`
}
