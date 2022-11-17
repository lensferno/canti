package network

import (
	"canti/app/api"
	"canti/app/codecs"
	"canti/app/commons"
	"canti/app/model"
	"encoding/json"
	"github.com/go-resty/resty/v2"
)

var defaultHeaders = map[string]string{
	"Accept":           "text/javascript, application/javascript, application/ecmascript, application/x-ecmascript, */*; q=0.01",
	"Accept-Encoding":  "gzip, deflate",
	"Accept-Language":  "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6",
	"Cache-Control":    "no-cache",
	"Connection":       "keep-alive",
	"Host":             "59.68.177.183",
	"Pragma":           "",
	"Referer":          "http://59.68.177.183/srun_portal_pc?ac_id=7&theme=pro",
	"User-Agent":       "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/107.0.0.0 Safari/537.36 Edg/107.0.1418.42",
	"X-Requested-With": "XMLHttpRequest",
}

func sendGetRequest(url string, param QueryParam) (string, error) {
	client := resty.New()
	resp, err := client.R().
		SetQueryParams(param).
		SetHeaders(defaultHeaders).
		Get(url)

	if err != nil {
		return "", err
	}

	return commons.FilterJQueryPrefix(resp.String()), nil
}

func RequestAuth(username, password, ip, challengeCode string) (string, error) {
	md5Password := codecs.HmacMd5(challengeCode, password)

	jsonMap := infoBody{
		Username: username,
		Password: password,
		Ip:       ip,
		Acid:     "7",
		EncVer:   "srun_bx1",
	}
	infoData, _ := json.Marshal(jsonMap)
	info := string(infoData)

	encodedInfo := `{SRBX1}` + codecs.SRBX1Encode(info, challengeCode)

	checksum := codecs.Checksum(challengeCode, username, md5Password, ip, encodedInfo)

	return sendGetRequest(api.AuthApi, authQuery(username, `{md5}`+md5Password, checksum, encodedInfo, ip))
}

type infoBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Ip       string `json:"ip"`
	Acid     string `json:"acid"`
	EncVer   string `json:"enc_ver"`
}

func RequestChallengeCode(username, ip string) (string, error) {
	return sendGetRequest(api.ChallengeCodeApi, challengeCodeQuery(username, ip))
}

func RequestUserInfo() (string, error) {
	return sendGetRequest(api.UserInfoApi, userInfoQuery())
}

func GetClientIp() (ip string, errs error) {
	userInfoJson, err := RequestUserInfo()
	if err != nil {
		return "", err
	}

	var userInfo model.UserInfo
	err2 := json.Unmarshal([]byte(userInfoJson), &userInfo)
	if err2 != nil {
		return "", err2
	}

	return userInfo.OnlineIp, nil
}
