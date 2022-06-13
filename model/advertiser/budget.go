package advertiser

// Budget 广告主预算
type Budget struct {
	// DayBudget 单日预算 单位：厘
	DayBudget int64 `json:"day_budget,omitempty"`
	// DayBudgetSchedule  分日预算; 单位：厘，单日预算和分日预算同时存在时，以分日预算为准
	DayBudgetSchedule []int64 `json:"day_budget_schedule,omitempty"`
}
