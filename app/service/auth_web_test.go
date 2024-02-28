package service

import (
	"canti/app/conf"
	"fmt"
	"os"
	"regexp"
	"testing"
)

func TestGetAnnounce(t *testing.T) {
	srv := NewService(conf.Config{})
	str, err := srv.GetWebAnnounce()
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Println(str)
	fmt.Println("----------------")
	regex := regexp.MustCompile("<.*?>(.*?)</.*?>")
	fmt.Printf("%s\n", regex.ReplaceAll([]byte(str), []byte("$1")))
}

func TestGetIp(t *testing.T) {
	srv := NewService(conf.Config{})
	ip, err := srv.getAuthIpAddress()
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Printf("ip: '%s'\n", ip)
}

func TestWebLogin(t *testing.T) {
	username, usernameProvided := os.LookupEnv("CANTI_TEST_USERNAME")
	password, passwordProvided := os.LookupEnv("CANTI_TEST_PASSWORD")

	if !usernameProvided || !passwordProvided {
		t.Error(fmt.Errorf("username and password wasn't provided in WebLogin test"))
	}

	config := conf.Config{
		Username:  username,
		Password:  password,
		Method:    "web",
		Reconnect: true,
		Silence:   false,
	}

	fmt.Printf("%+v\n", config)

	srv := NewService(config)
	userInfo, err := srv.WebLogin()
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Printf("Userinfo: %+v\n", userInfo)
}

func TestLogout(t *testing.T) {
	srv := NewService(conf.Config{})
	err := srv.WebLogout()
	if err != nil {
		t.Error(err)
		return
	}
}
