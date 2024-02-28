package service

import (
	"canti/library/ecode"
	"fmt"
	"regexp"
	"time"
)

const (
	CodeOk             = 0
	CodeOffline        = 1
	CodeRequireCaptcha = 2
)

const (
	webAuthBaseUrl = "http://59.68.177.9"
	apiPath        = "/api"

	_indexRedirectApiUrl = webAuthBaseUrl + apiPath + "/r/2"
	_configApiUrl        = webAuthBaseUrl + apiPath + "/config"
	_announceApiUrl      = webAuthBaseUrl + apiPath + "/notice/login"
	_accountStatusApiUrl = webAuthBaseUrl + apiPath + "/account/status"
	_accountCheckApiUrl  = webAuthBaseUrl + apiPath + "/account/check"
	_loginApiUrl         = webAuthBaseUrl + apiPath + "/account/login"
	_logoutApiUrl        = webAuthBaseUrl + apiPath + "/account/logout"
)

type GeneralResp[T any] struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data T      `json:"data"`
}

func (s *Service) GetWebAnnounce() (string, error) {
	req := s.requester.R()
	respJson := GeneralResp[string]{}
	resp, err := req.SetResult(&respJson).Get(_announceApiUrl)
	if err != nil {
		return "", err
	} else if !resp.IsSuccess() {
		return "", nil
	}

	if respJson.Code == CodeOk {
		return respJson.Data, nil
	} else {
		return respJson.Msg, nil
	}
}

type WebAccountCheckResp struct {
	Code        int    `json:"code"`
	IsChangePwd string `json:"isChangePwd"`
	Msg         string `json:"msg"`
}

type WebLoginResp struct {
	AuthCode   string `json:"authCode"`
	AuthMsg    string `json:"authMsg"`
	Code       int    `json:"code"`
	DialCode   string `json:"dialCode"`
	DialMsg    string `json:"dialMsg"`
	EnableDial bool   `json:"enableDial"`
	Msg        string `json:"msg"`
	Online     struct {
		AddTime        time.Time `json:"AddTime"`
		BytesIn4       string    `json:"BytesIn4"`
		Name           string    `json:"Name"`
		UserIpv4       string    `json:"UserIpv4"`
		UserMac        string    `json:"UserMac"`
		UserSourceType string    `json:"UserSourceType"`
		Username       string    `json:"Username"`
	} `json:"online"`
}

type OnlineStatus struct {
	Time           time.Time `json:"Time"`           // 登陆时间
	Bytes          string    `json:"Bytes"`          // 已用流量（字节）
	Name           string    `json:"Name"`           // 用户名（姓名）
	Ip             string    `json:"Ip"`             // ip地址
	Mac            string    `json:"Mac"`            // mac地址
	Username       string    `json:"Username"`       // 用户名（账号，即学号）
	UserSourceType string    `json:"UserSourceType"` // 用户来源类型
}

func (s *Service) getAuthIpAddress() (string, error) {
	resp, err := s.requester.R().Get(_indexRedirectApiUrl)
	if err != nil {
		return "", ecode.RequestErr
	}

	redirect := resp.Header().Get("Location")
	regex := regexp.MustCompile(".*?ip=(\\d+.\\d+.\\d+.\\d+).*")
	return regex.ReplaceAllString(redirect, "$1"), nil
}

const (
	_webLoginRefererFormat = "%s/tpl/wust/login.html?ip=%s&nasId=2"
)

func (s *Service) WebLogin() (*OnlineStatus, error) {
	ip, err := s.getAuthIpAddress()
	if err != nil {
		return nil, err
	}

	additionalHeaders := map[string]string{
		"Origin":  webAuthBaseUrl,
		"Referer": fmt.Sprintf(_webLoginRefererFormat, webAuthBaseUrl, ip),
	}

	webAccountCheckResp := new(WebAccountCheckResp)
	resp, err := s.requester.R().
		SetHeaders(additionalHeaders).
		SetResult(webAccountCheckResp).
		SetFormData(map[string]string{
			"username": s.conf.Username, "password": s.conf.Password,
			// nasId=2时貌似可以不用验证码
			"nasId": "2",
		}).
		Post(_accountCheckApiUrl)
	if err != nil {
		return nil, ecode.RequestErr
	} else if !resp.IsSuccess() {
		return nil, ecode.RequestErr
	}

	if webAccountCheckResp.Code != CodeOk {
		return nil, ecode.NewErrCode(webAccountCheckResp.Code, webAccountCheckResp.Msg)
	}

	webLoginResp := new(WebLoginResp)
	resp, err = s.requester.R().
		SetHeaders(additionalHeaders).
		SetResult(webLoginResp).
		SetFormData(map[string]string{
			"username": s.conf.Username, "password": s.conf.Password,
			// nasId=2时貌似可以不用验证码
			"nasId": "2",
		}).
		Post(_loginApiUrl)
	if err != nil {
		return nil, ecode.RequestErr
	} else if !resp.IsSuccess() {
		return nil, ecode.RequestErr
	}

	if webLoginResp.Code != CodeOk {
		return nil, ecode.NewErrCode(webLoginResp.Code, webLoginResp.Msg)
	}

	statusResp := webLoginResp.Online
	onlineStatus := OnlineStatus{
		Time:           statusResp.AddTime,
		Bytes:          statusResp.BytesIn4,
		Name:           statusResp.Name,
		Ip:             statusResp.UserIpv4,
		Mac:            statusResp.UserMac,
		UserSourceType: statusResp.UserSourceType,
		Username:       statusResp.Username,
	}

	return &onlineStatus, nil
}

func (s *Service) WebLogout() error {
	logoutResp := new(GeneralResp[string])
	resp, err := s.requester.R().
		SetResult(logoutResp).
		Get(_logoutApiUrl)
	if err != nil {
		return ecode.RequestErr
	} else if !resp.IsSuccess() {
		return ecode.RequestErr
	}

	if logoutResp.Code != CodeOk {
		return ecode.NewErrCode(logoutResp.Code, logoutResp.Msg)
	}

	return nil
}

type WebOnlineStatusResp struct {
	AuthCode   string `json:"authCode"`
	AuthMsg    string `json:"authMsg"`
	Code       int    `json:"code"`
	DialCode   string `json:"dialCode"`
	DialMsg    string `json:"dialMsg"`
	EnableDial bool   `json:"enableDial"`
	Msg        string `json:"msg"`
	Online     struct {
		AddTime        time.Time `json:"AddTime"`
		BytesIn4       string    `json:"BytesIn4"`
		Name           string    `json:"Name"`
		UserIpv4       string    `json:"UserIpv4"`
		UserMac        string    `json:"UserMac"`
		UserSourceType string    `json:"UserSourceType"`
		Username       string    `json:"Username"`
	} `json:"online"`
}

func (s *Service) WebGetOnlineStatus() (*OnlineStatus, error) {
	req := s.requester.R()

	webOnlineStatusResp := new(WebOnlineStatusResp)
	resp, err := req.SetResult(webOnlineStatusResp).Get(_accountStatusApiUrl)
	if err != nil {
		return nil, ecode.RequestErr
	} else if !resp.IsSuccess() {
		return nil, ecode.RequestErr
	}

	if webOnlineStatusResp.Code != CodeOk {
		return nil, ecode.NewErrCode(webOnlineStatusResp.Code, webOnlineStatusResp.Msg)
	}

	statusResp := webOnlineStatusResp.Online
	onlineStatus := OnlineStatus{
		Time:           statusResp.AddTime,
		Bytes:          statusResp.BytesIn4,
		Name:           statusResp.Name,
		Ip:             statusResp.UserIpv4,
		Mac:            statusResp.UserMac,
		UserSourceType: statusResp.UserSourceType,
		Username:       statusResp.Username,
	}

	return &onlineStatus, nil
}
