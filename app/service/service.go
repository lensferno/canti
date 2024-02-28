package service

import (
	"canti/app/conf"
	"github.com/go-resty/resty/v2"
	"net/http"
	"time"
)

type Service struct {
	conf      *conf.Config
	requester *resty.Client
}

var _defaultHeaders = map[string]string{
	"Accept":           "*/*",
	"Accept-Encoding":  "gzip, deflate",
	"Accept-Language":  "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6",
	"Cache-Control":    "no-cache",
	"Pragma":           "no-cache",
	"User-Agent":       "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.0.0 Safari/537.36 Edg/121.0.0.0",
	"X-Requested-With": "XMLHttpRequest",
}

func NewService(config conf.Config) *Service {
	requester := resty.New().SetHeaders(_defaultHeaders)
	requester.GetClient().CheckRedirect = func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	}
	requester.SetTimeout(10 * time.Second)

	return &Service{
		conf:      &config,
		requester: requester,
	}
}

func (s *Service) SetConfig(config conf.Config) *Service {
	*s.conf = config
	return s
}
