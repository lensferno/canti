package main

var helpTemplate = `
 ██████╗ █████╗ ███╗   ██╗████████╗██╗
██╔════╝██╔══██╗████╗  ██║╚══██╔══╝██║
██║     ███████║██╔██╗ ██║   ██║   ██║
██║     ██╔══██║██║╚██╗██║   ██║   ██║
╚██████╗██║  ██║██║ ╚████║   ██║   ██║
 ╚═════╝╚═╝  ╚═╝╚═╝  ╚═══╝   ╚═╝   ╚═╝

武科大wust校园网认证登录（web + pppoe拨号）
(当然啦，是第三方的)
---------------------------------------
{{if .VisibleCommands}}
命令:{{template "visibleCommandCategoryTemplate" .}}{{end}}
{{if .VisibleFlagCategories}}
选项:{{template "visibleFlagCategoryTemplate" .}}
{{else if .VisibleFlags}}
选项:{{template "visibleFlagTemplate" .}}{{end}}

如：
    1. 使用web方法认证，用户名（学号）为202100000000，密码为12450password：
    ./canti --username 202100000000 --password 12450password -m web

    2. 安装为开机自启的服务（用户信息同上文，需要管理员权限）：
    ./canti install --username 202100000000 --password 12450password -m web

---------------------------------------
{{.Version}}
made by @lensfrex

本项目在Github上开源：(GPLv3)
https://github.com/lensferno/canti
（详细的文档也在上面哦）
如果有什么好的建议之类的请务必告诉我(。・∀・)ノ
`
