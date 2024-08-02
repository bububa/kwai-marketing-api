package model

import (
	"encoding/json"
	"strconv"
)

// Uint64 support string quoted number in json
type Uint64 uint64

// UnmarshalJSON implement json Unmarshal interface
func (u64 *Uint64) UnmarshalJSON(b []byte) (err error) {
	if len(b) == 0 {
		return nil
	}
	if b[0] == '"' && b[len(b)-1] == '"' {
		b = b[1 : len(b)-1]
	}
	i, _ := strconv.ParseUint(string(b), 10, 64)
	*u64 = Uint64(i)
	return
}

func (u64 Uint64) Value() uint64 {
	return uint64(u64)
}

func (u64 Uint64) String() string {
	return strconv.FormatUint(uint64(u64), 10)
}

// Int64 support string quoted number in json
type Int64 int64

// UnmarshalJSON implement json Unmarshal interface
func (i64 *Int64) UnmarshalJSON(b []byte) (err error) {
	if len(b) == 0 {
		return nil
	}
	if b[0] == '"' && b[len(b)-1] == '"' {
		b = b[1 : len(b)-1]
	}
	i, _ := strconv.ParseInt(string(b), 10, 64)
	*i64 = Int64(i)
	return
}

func (i64 Int64) Value() int64 {
	return int64(i64)
}

func (i64 Int64) String() string {
	return strconv.FormatInt(int64(i64), 10)
}

// Int support string quoted number in json
type Int int

// UnmarshalJSON implement json Unmarshal interface
func (i *Int) UnmarshalJSON(b []byte) (err error) {
	if len(b) == 0 {
		return nil
	}
	if b[0] == '"' && b[len(b)-1] == '"' {
		b = b[1 : len(b)-1]
	}
	v, _ := strconv.Atoi(string(b))
	*i = Int(v)
	return
}

func (i Int) Value() int {
	return int(i)
}

func (i Int) String() string {
	return strconv.Itoa(int(i))
}

// Float64 support string quoted number in json
type Float64 float64

// UnmarshalJSON implement json Unmarshal interface
func (f64 *Float64) UnmarshalJSON(b []byte) (err error) {
	if len(b) == 0 {
		return nil
	}
	if b[0] == '"' && b[len(b)-1] == '"' {
		b = b[1 : len(b)-1]
	}
	i, _ := strconv.ParseFloat(string(b), 64)
	*f64 = Float64(i)
	return
}

func (f64 Float64) Value() float64 {
	return float64(f64)
}

func (f64 Float64) String(prec int) string {
	return strconv.FormatFloat(float64(64), 'f', prec, 64)
}

func JSONMarshal(i interface{}) []byte {
	b, _ := json.Marshal(i)
	return b
}

type MatchType int

// UnmarshalJSON implement json Unmarshal interface
func (m *MatchType) UnmarshalJSON(b []byte) (err error) {
	if len(b) == 0 {
		return nil
	}
	if b[0] == '"' && b[len(b)-1] == '"' {
		b = b[1 : len(b)-1]
	}
	if i, err := strconv.ParseFloat(string(b), 64); err == nil {
		*m = MatchType(i)
	} else {
		switch string(b) {
		case "精确匹配":
			*m = MatchType(1)
		case "短语匹配":
			*m = MatchType(2)
		case "广泛匹配":
			*m = MatchType(3)
		}
	}
	return
}

// UnmarshalCSV implement json Unmarshal interface
func (m *MatchType) UnmarshalCSV(b string) (err error) {
	if len(b) == 0 {
		return nil
	}
	if b[0] == '"' && b[len(b)-1] == '"' {
		b = b[1 : len(b)-1]
	}
	if i, err := strconv.ParseFloat(b, 64); err == nil {
		*m = MatchType(i)
	} else {
		switch b {
		case "精确匹配":
			*m = MatchType(1)
		case "短语匹配":
			*m = MatchType(2)
		case "广泛匹配":
			*m = MatchType(3)
		}
	}
	return
}

func (m MatchType) Value() int {
	return int(m)
}

// PageInfo 分页信息
type PageInfo struct {
	// PageSize 页大小
	PageSize int `json:"page_size,omitempty"`
	// CurrentPage 当前页码
	CurrentPage int `json:"current_page,omitempty"`
	// TotalCount 总记录数
	TotalCount int `json:"total_count,omitempty"`
}
