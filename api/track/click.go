package track

import (
	"net/url"
)

var DEFAULT_CLICK_FIELDS = []string{
	"aid",
	"advertiser_id",
	"cid",
	"campaign_id",
	"csite",
	"imei",
	"idfa",
	"android_id",
	"oaid",
	"os",
	"ip",
	"ua",
	"ts",
	"callback",
	"model",
}

// 点击检测链接
func Click(baseUrl string, fields []string) string {
	if fields == nil {
		fields = DEFAULT_CLICK_FIELDS
	}
	parsedUrl, _ := url.Parse(baseUrl)
	values := parsedUrl.Query()
	for _, field := range fields {
		switch field {
		case "advertiser_id":
			values.Set("advertiser_id", "ACCOUNTID")
		case "aid":
			values.Set("aid", "AID")
		case "cid":
			values.Set("cid", "CID")
		case "campaign_id":
			values.Set("campaign_id", "DID")
		case "campaign_name":
			values.Set("campaign_name", "DNAM")
		case "csite":
			values.Set("csite", "CSITE")
		case "imei":
			values.Set("imei", "IMEI2")
		case "idfa":
			values.Set("idfa", "IDFA2")
		case "android_id":
			values.Set("android_id", "ANDROIDID2")
		case "oaid":
			values.Set("oaid", "OAID")
		case "os":
			values.Set("os", "OS")
		case "ip":
			values.Set("ip", "IP")
		case "ua":
			values.Set("ua", "UA")
		case "ts":
			values.Set("ts", "TS")
		case "callback":
			values.Set("callback", "CALLBACK")
		case "model":
			values.Set("model", "MODEL")
		}
	}
	parsedUrl.RawQuery = values.Encode()
	return parsedUrl.String()
}
