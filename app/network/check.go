package network

import (
	"canti/app/api"
	"github.com/go-resty/resty/v2"
)

const (
	TestBaidu         = "http://baidu.com"
	TestTencent       = "http://qq.com"
	TestTaobao        = "http://taobao.com"
	TestSchoolNetwork = api.Host
)

// CheckNetwork 检查网络是否可通
func CheckNetwork(testUrl string) bool {
	client := resty.New()
	resp, err := client.R().SetHeaders(defaultHeaders).Get(testUrl)

	return err != nil && resp.Body() != nil
}
