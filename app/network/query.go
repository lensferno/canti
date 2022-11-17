package network

import (
	"canti/app/commons"
	"strconv"
	"time"
)

type QueryParam map[string]string

func (param QueryParam) Add(key string, value string) {
	param[key] = value
}

// generateCallback callback参数生成，虽然说实际上callback参数随意设置也行，但是还是生成一个类似的比较好
func generateCallback() string {
	return "jQuery_1124005588867363182781_" + strconv.FormatInt(time.Now().UnixMilli(), 10)
}

func userInfoQuery() QueryParam {
	return QueryParam{
		"callback": generateCallback(),
	}
}

func challengeCodeQuery(username string, ip string) QueryParam {
	return QueryParam{
		"callback": generateCallback(),
		"username": username,
		"ip":       ip,
		"_":        commons.CurrentMilliSecond(),
	}
}

func authQuery(username, password, checkSum, info, ip string) QueryParam {
	return QueryParam{
		"callback":     generateCallback(),
		"action":       "login",
		"username":     username,
		"password":     password,
		"os":           commons.GetOsType(),
		"name":         commons.GetOsName(),
		"double_stack": "0",
		"chksum":       checkSum,
		"info":         info,
		"ac_id":        "7",
		"ip":           ip,
		"n":            "200",
		"type":         "1",
		"_":            commons.CurrentMilliSecond(),
	}
}
