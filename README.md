# Canti

> 仍在施工中🚧
> 
> 有好的建议和意见欢迎在issue里提出~

全自动的wust武科大校园网认证客户端（当然，是第三方的）

支持使用web方式认证和pppoe拨号（尚未实现）

使用Go语言编写，支持多个平台使用

# 使用

- 使用web方法认证，用户名（学号）为202100000000，密码为12450password：
  ``` bash
  ./canti login --username 202100000000 --password 12450password -m web
  ```

- 按照指定的配置进行登录认证：
  ``` bash
  ./canti login --config ./config.yml
  ```
  支持的文件格式有：JSON, TOML, YAML, HCL, .env

- 安装为开机自启的服务（用户信息同上文，需要管理员权限）：
  ``` bash
  ./canti install
  ```
	按照提示填写信息即可

运行./canti [命令] --help获取相关参数