package model

// TemporalGranularityType 时间粒度
type TemporalGranularityType string

const (
	// TemporalGranularityType_DAILY 天粒度
	TemporalGranularityType_DAILY TemporalGranularityType = "DAILY"
	// TemporalGranularityType_HOURLY 小时粒度
	TemporalGranularityType_HOURLY TemporalGranularityType = "HOURLY"
)
