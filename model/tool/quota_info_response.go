package tool

type QuotaInfoResponse struct {
	QuotaDefinition string `json:"quota_definition"`
	IsInQuotaScope  bool   `json:"is_in_quota_scope"`
	AvailableCount  int    `json:"available_count"`
	QuotaCount      int    `json:"quota_count"`
	QuotaMsgList    []struct {
		Msg  string `json:"msg"`
		Type string `json:"type"`
	} `json:"quota_msg_list"`
}