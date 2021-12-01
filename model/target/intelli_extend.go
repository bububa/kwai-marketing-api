package target

// IntelliExtend 智能扩量
type IntelliExtend struct {
	// IsOpen 开启智能扩量;0：关闭智能扩量，1：打开智能扩量
	IsOpen int `json:"is_open"`
	// NoAgeBreak 不可突破年龄;0：可突破/无需控制，1：不可突破
	NoAgeBreak int `json:"no_age_break"`
	// NoGenderBreak 不可突破性别;0：可突破/无需控制，1：不可突破
	NoGenderBreak int `json:"no_gender_break"`
	// NoAreaBreak 不可突破地域; 0：可突破/无需控制，1：不可突破
	NoAreaBreak int `json:"no_area_break"`
}
