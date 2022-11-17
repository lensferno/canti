package main

import (
	"fmt"
	"github.com/kardianos/service"
	"github.com/urfave/cli/v2"
	"os"
)

var serviceConfig = &service.Config{
	Name:        "Canti",
	DisplayName: "Canti",
	Description: "武科大wust校园网自动认证登录服务",
}

func install() {
	service, err := service.New(nil, serviceConfig)
	if err != nil {
		err2 := fmt.Errorf("安装服务时出错： %s", err)
		fmt.Println(err2)

		return
	}

	serviceConfig.Arguments = os.Args[2:]

	serviceInstallError := service.Install()
	if serviceInstallError != nil {
		err2 := fmt.Errorf("安装服务时出错： %s", serviceInstallError)
		fmt.Println(err2)

		return
	}
}

func installService(ctx *cli.Context) error {

	return nil
}

func uninstallService(ctx *cli.Context) error {
	return nil
}

func login(ctx *cli.Context) error {
	return nil
}

func logout(ctx *cli.Context) error {
	return nil
}

func showStatus(ctx *cli.Context) error {
	return nil
}
