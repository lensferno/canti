package main

import "github.com/urfave/cli/v2"

var globalFlags = []cli.Flag{
	&cli.StringFlag{
		Name:        "config",
		Aliases:     []string{"c"},
		Usage:       "指定`file/path`为配置文件，如果指定了有效的配置文件，则命令行参数不会生效，仅使用配置文件提供的参数",
		Destination: nil,
	},
	&cli.StringFlag{
		Name:        "username",
		Aliases:     []string{"u"},
		Usage:       "校园网登录用户名（账号）",
		Destination: nil,
	},
	&cli.StringFlag{
		Name:        "password",
		Aliases:     []string{"p"},
		Usage:       "账号的密码",
		Destination: nil,
	},
	&cli.StringFlag{
		Name:        "method",
		Aliases:     []string{"m"},
		Value:       "web",
		Usage:       "认证方法，可选的值为web认证（仅无线网络用户）或pppoe认证（仅有线网络用户），默认使用web认证，当一种失败后会自动切换另外一种方式重试",
		Destination: nil,
	},
	&cli.BoolFlag{
		Name:        "keep-alive",
		Aliases:     []string{"k"},
		Value:       false,
		Usage:       "是否定时（1分钟）请求一次网络以防止自动断线，默认不开启，开启后，程序将不会自动退出",
		Destination: nil,
	},
	&cli.BoolFlag{
		Name:    "reconnect",
		Aliases: []string{"r"},
		Value:   true,
		Usage:   "掉线时是是否重新登录，默认开启，开启后，程序将不会自动退出",
	},
}
