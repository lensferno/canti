package service

import (
	"canti/app/model"
	"canti/app/network"
	"encoding/json"
	"errors"
	"fmt"
)

var NoNeedLogin = errors.New("no need to login")
var UserinfoError = errors.New("error happens in requesting userinfo")
var ChallengeCodeError = errors.New("error happens in requesting challenge code")

func WebLogin(username string, password string) error {
	// 能直接上网的话就不用登陆了
	if network.CheckNetwork(network.TestBaidu) {
		return NoNeedLogin
	}

	ip, err := network.GetClientIp()
	if err != nil || ip == "" || ip == "<nil>" {
		return UserinfoError
	}

	fmt.Print("Got ip:")
	fmt.Println(ip)

	challengeCodeJson, err2 := network.RequestChallengeCode(username, ip)
	if err2 != nil || challengeCodeJson == "" {
		return err
	}

	var challengeCodeResponse model.ChallengeCodeResponse
	jsonError := json.Unmarshal([]byte(challengeCodeJson), &challengeCodeResponse)
	if jsonError != nil {
		return jsonError
	} else if challengeCodeResponse.Challenge == "" {
		_ = fmt.Errorf(challengeCodeJson)
		return ChallengeCodeError
	}

	fmt.Print(challengeCodeResponse)

	return nil
}
