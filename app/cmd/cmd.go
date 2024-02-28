package cmd

import (
	"canti/app/conf"
	"canti/app/service"
	"canti/app/version"
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"regexp"
)

var (
	srv *service.Service
)

func Start() {
	cli.AppHelpTemplate = helpTemplate

	app := &cli.App{
		//Flags:    globalFlags,
		Commands: allCommands,
		Version:  version.Version,
	}

	srv = service.NewService(conf.Config{
		Username:  "",
		Password:  "",
		Method:    "web",
		Reconnect: true,
		Silence:   false,
	})
	announce, err := srv.GetWebAnnounce()
	if err != nil {
		fmt.Printf("获取公告时出现错误：%s", err.Error())
	} else {
		fmt.Println("--------------------------------")
		fmt.Println("来自系统的公告：")
		fmt.Println("--------------------------------")
		regex := regexp.MustCompile("<.*?>(.*?)</.*?>")
		fmt.Println(regex.ReplaceAllString(announce, "$1"))
		fmt.Println("--------------------------------")
	}

	err = app.Run(os.Args)
	if err != nil {
		log.Fatalln(err)
		return
	}
}
