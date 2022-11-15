package main

import (
	"fmt"
	"github.com/kardianos/service"
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
