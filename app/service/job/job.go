package job

import (
	"canti/app/service"
	"fmt"
	"github.com/mgutz/ansi"
	"time"
)

type LoginStatusKeepingJob struct {
	srv *service.Service
}

func NewLoginStatusKeepingJob(srv *service.Service) *LoginStatusKeepingJob {
	return &LoginStatusKeepingJob{
		srv: srv,
	}
}

func (l *LoginStatusKeepingJob) Run() {
	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()

	for range ticker.C {
		networkAvailable := l.srv.TestNetwork()
		if networkAvailable {
			continue
		}

		fmt.Println(ansi.Color("网络疑似断开，尝试重新登录", ansi.Yellow))
		onlineStatus, err := l.srv.WebLogin()
		if err != nil {
			fmt.Printf(ansi.Color("错误：%s\n", ansi.Yellow), err.Error())
			continue
		}

		fmt.Printf("%+v\n", *onlineStatus)
	}
}
