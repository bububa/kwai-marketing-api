package track

import (
	"net/url"
)

var DEFAULT_CLICK_FIELDS = []string{
	"request_id",
	"aid",
	"advertiser_id",
	"cid",
	"campaign_id",
	"csite",
	"imei",
	"idfa",
	"android_id",
	"oaid",
	"caid",
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
		case "request_id":
			values.Set("request_id", "__LLSID__")
		case "advertiser_id":
			values.Set("advertiser_id", "__ACCOUNTID__")
		case "aid":
			values.Set("aid", "__AID__")
		case "cid":
			values.Set("cid", "__CID__")
		case "campaign_id":
			values.Set("campaign_id", "__DID__")
		case "campaign_name":
			values.Set("campaign_name", "__DNAM__")
		case "csite":
			values.Set("csite", "__CSITE__")
		case "imei":
			values.Set("imei", "__IMEI4__")
		case "idfa":
			values.Set("idfa", "__IDFA2__")
		case "android_id":
			values.Set("android_id", "__ANDROIDID2__")
		case "oaid":
			values.Set("oaid", "__OAID2__")
		case "caid":
			values.Set("caid", "__KENYID_CAA__")
		case "os":
			values.Set("os", "OS")
		case "ip":
			values.Set("ip", "__IP__")
		case "ua":
			values.Set("ua", "__UA__")
		case "ts":
			values.Set("ts", "__TS__")
		case "callback":
			values.Set("callback", "__CALLBACK__")
		case "model":
			values.Set("model", "__MODEL__")
		}
	}
	parsedUrl.RawQuery = values.Encode()
	return parsedUrl.String()
}
