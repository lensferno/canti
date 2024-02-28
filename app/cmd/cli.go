package cmd

import (
	"github.com/urfave/cli/v2"
)

var allCommands = []*cli.Command{
	{
		Name:    "install",
		Aliases: []string{"i"},
		Usage:   "将canti安装为服务以支持开机自启（需要管理员权限）",
		Action:  createService,
	},
	{
		Name:    "uninstall",
		Aliases: []string{"u"},
		Usage:   "卸载canti的服务（需要管理员权限）",
		Action:  removeService,
	},
	{
		Name:    "login",
		Aliases: []string{"l"},
		Flags:   loginFlags,
		Usage:   "登录校园网",
		Action:  login,
	},
	{
		Name:   "logout",
		Usage:  "退出登录状态",
		Action: logout,
	},
	{
		Name:    "status",
		Aliases: []string{"info"},
		Usage:   "查看当前登陆状态",
		Action:  status,
	},
}

type Config struct {
	ConfigFile string `json:"configFile" yaml:"configFile"`
	Username   string `json:"username" yaml:"username"`
	Password   string `json:"password" yaml:"password"`
	Method     string `json:"method" yaml:"method"`
	Reconnect  bool   `json:"reconnect" yaml:"reconnect"`
	Silence    bool   `json:"silence" yaml:"silence"`
}

var globalConfig = &Config{}

var loginFlags = []cli.Flag{
	&cli.StringFlag{
		Name:        "config",
		Aliases:     []string{"c"},
		Value:       "./config",
		Usage:       "指定配置文件或者包含配置文件的的目录（此时文件名需要为\"config\"(不包括json,yaml等扩展名)），如果指定了该配置项，则其他的命令行参数不会生效，而是使用配置文件提供的参数。支持的配置文件格式：JSON, TOML, YAML, HCL, .env",
		Destination: &globalConfig.ConfigFile,
	},
	&cli.StringFlag{
		Name:        "username",
		Aliases:     []string{"u"},
		Value:       "",
		Usage:       "校园网登录用户名（账号）",
		Destination: &globalConfig.Username,
	},
	&cli.StringFlag{
		Name:        "password",
		Aliases:     []string{"p"},
		Value:       "",
		Usage:       "账号的密码",
		Destination: &globalConfig.Password,
	},
	&cli.StringFlag{
		Name:        "method",
		Aliases:     []string{"m"},
		Value:       "web",
		Usage:       "认证方法，可选的值为web认证（仅无线网络用户）或pppoe认证（仅有线网络用户，功能尚未实现），默认使用web认证，当一种失败后会自动切换另外一种方式重试（当前仅支持web认证）",
		Destination: &globalConfig.Method,
	},
	&cli.BoolFlag{
		Name:        "reconnect",
		Aliases:     []string{"r"},
		Value:       true,
		Usage:       "是否定时（1分钟）请求一次网络以防止自动断线，并且掉线时是是否重新登录，默认开启，开启后，程序将不会自动退出",
		Destination: &globalConfig.Reconnect,
	},
	&cli.BoolFlag{
		Name:        "silence",
		Aliases:     []string{"s"},
		Value:       false,
		Usage:       "是否静默执行，设置为true时，仅会在出现错误时输出信息",
		Destination: &globalConfig.Silence,
	},
}

var helpTemplate = `
 ██████╗ █████╗ ███╗   ██╗████████╗██╗
██╔════╝██╔══██╗████╗  ██║╚══██╔══╝██║
██║     ███████║██╔██╗ ██║   ██║   ██║
██║     ██╔══██║██║╚██╗██║   ██║   ██║
╚██████╗██║  ██║██║ ╚████║   ██║   ██║
 ╚═════╝╚═╝  ╚═╝╚═╝  ╚═══╝   ╚═╝   ╚═╝

全自动的wust武科大校园网认证客户端（web + pppoe拨号）
(当然，是第三方的)
---------------------------------------
{{if .VisibleCommands}}
命令:{{template "visibleCommandCategoryTemplate" .}}{{end}}
{{if .VisibleFlagCategories}}
选项:{{template "visibleFlagCategoryTemplate" .}}
{{else if .VisibleFlags}}
选项:{{template "visibleFlagTemplate" .}}{{end}}

如：
	0. 使用web方法认证，用户名（学号）为202100000000，密码为12450password：
    ./canti login --username 202100000000 --password 12450password -m web

    1. 按照指定的配置进行登录认证：
    ./canti login --config ./config.yml
	其中支持的文件格式有：JSON, TOML, YAML, HCL, .env

    2. 安装为开机自启的服务（用户信息同上文，需要管理员权限）：
    ./canti install
	按照提示填写信息即可

运行./canti [命令] --help获取相关参数

---------------------------------------
{{.Version}}
made by @lensfrex, lensferno@outlook.com

本项目在Github上开源(GPLv3)，代码完全开放：
https://github.com/lensferno/canti

如果有什么好的建议之类的请务必告诉我(。・∀・)ノ
`
