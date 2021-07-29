# 快手磁力引擎MarketingAPI Golang SDK

[![Go Reference](https://pkg.go.dev/badge/github.com/bububa/kwai-marketing-api.svg)](https://pkg.go.dev/github.com/bububa/kwai-marketing-api)
[![Go](https://github.com/bububa/kwai-marketing-api/actions/workflows/go.yml/badge.svg)](https://github.com/bububa/kwai-marketing-api/actions/workflows/go.yml)
[![goreleaser](https://github.com/bububa/kwai-marketing-api/actions/workflows/goreleaser.yml/badge.svg)](https://github.com/bububa/kwai-marketing-api/actions/workflows/goreleaser.yml)
[![GitHub go.mod Go version of a Go module](https://img.shields.io/github/go-mod/go-version/bububa/kwai-marketing-api.svg)](https://github.com/bububa/kwai-marketing-api)
[![GoReportCard](https://goreportcard.com/badge/github.com/bububa/kwai-marketing-api)](https://goreportcard.com/report/github.com/bububa/kwai-marketing-api)
[![GitHub license](https://img.shields.io/github/license/bububa/kwai-marketing-api.svg)](https://github.com/bububa/kwai-marketing-api/blob/master/LICENSE)
[![GitHub release](https://img.shields.io/github/release/bububa/kwai-marketing-api.svg)](https://GitHub.com/bububa/kwai-marketing-api/releases/)


- Oauth2 授权 (api/oauth)
  - 生成授权链接 [ Url(clt *core.SDKClient, req *oauth.UrlRequest) string ]
  - 获取AccessToken [ AccessToken(clt *core.SDKClient, authCode String) (*oauth.AccessTokenResponse, error) ]
  - 刷新Token [ RefreshToken(clt *core.SDKClient, refreshToken string) (*oauth.AccessTokenResponse, error)]
- 账号服务
  - 广告主 (api/advertiser)
    - 获取广告主信息 [ Info(clt *core.SDKClient, accessToken string, advertiserID int64) (*advertiser.Info, error) ]
    - 获取广告账户余额信息 [ FundGet(clt *core.SDKClient, accessToken string, advertiserID int64) (float64, error) ]
    - 获取广告主账户流水信息 [ FundDailyFlows(clt *core.SDK, accessToken string, req *advertiser.FundDailyFlowsRequest) (*advertiser.FundDailyFlowsResponse, error) ] 
- 广告投放
  - 账户层级
    - 账户日预算查询 [ advertiser.BudgetGet(clt *core.SDKClient, accessToken string, advertiserID int64) (*advertiser.Budget, error) ]
    - 修改账户预算 [ advertiser.UpdateBudget(clt *core.SDKClient, accessToken string, req *advertiser.UpdateBudgetRequest) error ]
- 数据报表
  - 广告数据报表 (api/report)
    - 代理商数据 [ AgentReport(clt *core.SDKClient, accessToken string, req *report.AgentReportRequest) (*report.AgentReportResponse, error) ]
    - 广告主数据 [ AccountReport(clt *core.SDKClient, accessToken string, req *report.AccountReportRequest) (*report.ReportResponse, error) ]
    - 广告计划数据 [ CampaignReport(clt *core.SDKClient, accessToken string, req *report.CampaignReportRequest) (*report.ReportResponse, error) ]
    - 广告单元数据 [ UnitReport(clt *core.SDKClient, accessToken string, req *report.UnitReportRequest) (*report.ReportResponse, error) ]
    - 广告创意数据 [ CreativeReport(clt *core.SDKClient, accessToken string, req *report.CreativeReportRequest) (*report.ReportResponse, error) ]
    - 程序化创意数据 [ ProgramCreativeReport(clt *core.SDKClient, accessToken string, req *report.ProgramCreativeReportRequest) (*report.ReportResponse, error) ]
    - 广告素材数据 [ CreativeReport(clt *core.SDKClient, accessToken string, req *report.MaterialReportRequest) (*report.ReportResponse, error) ]
    - 人群分析数据 [ AudienceReport(clt *core.SDKClient, accessToken string, req *report.AudienceReportRequest) (*report.ReportResponse, error) ]
    - 小店通转化数据 [ MerchantDeatailReport(clt *core.SDKClient, accessToken string, req *report.MerchantDetailReportRequest) (*report.MerchantDetailReportResponse, error) ]
- 数据上报管理 (api/track)
  - 转化回传 [ Activate(req *track.ActivateRequest) error ]
