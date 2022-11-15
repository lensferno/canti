package main

import (
	"github.com/urfave/cli/v2"
)

var allCommands = []*cli.Command{
	{
		Name:    "install",
		Aliases: []string{"i"},
		Usage:   "将Canti安装为服务以支持开机自启（需要管理员权限）",
		Action:  installService,
	},
	{
		Name:    "uninstall",
		Aliases: []string{"u"},
		Usage:   "卸载Canti的服务（需要管理员权限）",
		Action:  uninstallService,
	},
	{
		Name:    "login",
		Aliases: []string{"l"},
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
		Aliases: []string{"a", "info"},
		Usage:   "查看当前登陆状态",
		Action:  showStatus,
	},
}
