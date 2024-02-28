package service

import "fmt"

var testUrls = []string{
	"https://www.baidu.com",
	"https://www.gov.cn",
	"https://www.qq.com/",
	"https://www.taobao.com/",
}

func (s *Service) TestNetwork() bool {
	for _, url := range testUrls {
		resp, err := s.requester.R().Get(url)
		if err == nil && resp.IsSuccess() {
			return true
		}

		fmt.Printf("network check: '%s' not available, err: %s, statusCode: %v , try next\n", url, err.Error(), resp.StatusCode())
	}

	return false
}
