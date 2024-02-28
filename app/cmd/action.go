package cmd

import (
	"canti/app/cmd/interaction"
	"canti/app/conf"
	"canti/app/service/job"
	"encoding/json"
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"github.com/kardianos/service"
	"github.com/mgutz/ansi"
	"github.com/spf13/viper"
	"github.com/urfave/cli/v2"
	"os"
	"path/filepath"
)

func createService(c *cli.Context) error {
	fmt.Println("安装canti作为服务")
	fmt.Println("在开始之前，需要获得一些信息")
	serviceInstallSurvey := interaction.NewServiceInstallSurvey()
	err := serviceInstallSurvey.Ask()
	if err != nil {
		fmt.Printf(ansi.Color("错误：%s\n", ansi.Red), err.Error())
		return err
	}

	config := &conf.Config{}
	config.Username = serviceInstallSurvey.Username
	config.Password = serviceInstallSurvey.Password
	config.Method = serviceInstallSurvey.Method
	config.Reconnect = serviceInstallSurvey.Reconnect
	config.Silence = false
	srv.SetConfig(*config)

	configFile := serviceInstallSurvey.Config
	fileInfo, err := os.Stat(configFile)
	if err != nil {
		fmt.Printf(ansi.Color("错误：%s\n", ansi.Red), err.Error())
		return err
	}

	if fileInfo.IsDir() {
		configFile = filepath.Join(configFile, "config.json")
	}

	configFile, err = filepath.Abs(configFile)
	if err != nil {
		fmt.Printf(ansi.Color("错误：%s\n", ansi.Red), err.Error())
		return err
	}

	fmt.Printf("配置文件将保存在：%s\n", configFile)

	jsonBytes, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		fmt.Printf(ansi.Color("错误：%s\n", ansi.Red), err.Error())
		return err
	}

	err = os.WriteFile(configFile, jsonBytes, 660)
	if err != nil {
		fmt.Printf(ansi.Color("错误：%s\n", ansi.Red), err.Error())
		return err
	}

	srvConfig := service.Config{
		Name:        "canti-auto-service",
		DisplayName: "canti",
		Description: "全自动的must武科大校园网认证客户端",
		Arguments:   []string{"login", "--config", configFile},
		Executable:  os.Args[0],
	}

	s, err := service.New(nil, &srvConfig)
	if err != nil {
		fmt.Printf(ansi.Color("错误：%s\n", ansi.Red), err.Error())
		return err
	}

	err = s.Install()
	if err != nil {
		fmt.Printf(ansi.Color("错误：%s\n", ansi.Red), err.Error())
		return err
	}

	tryLogin := false
	err = survey.AskOne(&survey.Confirm{
		Message: "好了，要试一试吗？（尝试前请先退出登录）：",
		Default: true,
	},
		&tryLogin, survey.WithValidator(survey.Required),
	)
	if err != nil {
		fmt.Printf(ansi.Color("错误：%s\n", ansi.Red), err.Error())
		return err
	}

	if !tryLogin {
		return nil
	}

	err = _loginWithConfig(config)
	if err != nil {
		fmt.Printf(ansi.Color("错误：%s\n", ansi.Red), err.Error())
		return err
	}

	return nil
}

func removeService(c *cli.Context) error {
	srvConfig := service.Config{
		Name: "canti-auto-service",
	}

	s, err := service.New(nil, &srvConfig)
	if err != nil {
		fmt.Printf(ansi.Color("错误：%s\n", ansi.Red), err.Error())
		return err
	}

	err = s.Uninstall()
	if err != nil {
		fmt.Printf(ansi.Color("错误：%s\n", ansi.Red), err.Error())
		return err
	}

	return nil
}

func login(c *cli.Context) error {
	confFile := c.String("config")
	config := conf.Config{}
	if confFile != "" {
		fileInfo, err := os.Stat(confFile)
		if err != nil {
			fmt.Printf(ansi.Color("错误：%s\n", ansi.Red), err.Error())
			return err
		} else {
			if !fileInfo.IsDir() {
				viper.SetConfigFile(confFile)
			} else {
				viper.AddConfigPath(confFile)
			}
		}

		err = viper.ReadInConfig()
		if err != nil {
			fmt.Printf(ansi.Color("错误：%s\n", ansi.Red), err.Error())
			return err
		}

		err = viper.Unmarshal(&config)
		if err != nil {
			fmt.Printf(ansi.Color("错误：%s\n", ansi.Red), err.Error())
			return err
		}
	} else {
		username := c.String("username")
		password := c.String("password")
		if username == "" || password == "" {
			fmt.Println(ansi.Color("错误：必须提供用户名和密码，或者指定配置文件", ansi.Red))
			return fmt.Errorf("必须提供用户名和密码，或者指定配置文件")
		}

		config.Username = username
		config.Password = password
		config.Method = _readCliParam(c, "method", conf.LoginWebMethod)
		config.Reconnect = _readCliParam(c, "reconnect", true)
		config.Silence = _readCliParam(c, "silence", false)
	}

	srv.SetConfig(config)

	if config.Reconnect {
		loginStatusKeepingJob := job.NewLoginStatusKeepingJob(srv)
		loginStatusKeepingJob.Run()
		return nil
	}

	return _loginWithConfig(&config)
}

func _loginWithConfig(config *conf.Config) error {
	srv.SetConfig(*config)
	onlineStatus, err := srv.WebLogin()
	if err != nil {
		fmt.Printf(ansi.Color("错误：%s\n", ansi.Red), err.Error())
		return err
	}

	fmt.Println(ansi.Color("登陆成功", ansi.Cyan))
	jsonBytes, _ := json.MarshalIndent(onlineStatus, "", "  ")
	fmt.Printf("用户信息：%s\n", string(jsonBytes))
	return nil
}

func logout(c *cli.Context) error {
	err := srv.WebLogout()
	if err != nil {
		fmt.Printf(ansi.Color("错误：%s\n", ansi.Red), err.Error())
		return err
	}

	return nil
}

func status(c *cli.Context) error {
	onlineStatus, err := srv.WebGetOnlineStatus()
	if err != nil {
		fmt.Printf(ansi.Color("错误：%s\n", ansi.Red), err.Error())
		return err
	}

	jsonBytes, _ := json.MarshalIndent(onlineStatus, "", "  ")
	//fmt.Println(ansi.Color("状态：在线", ansi.Blue))
	fmt.Println("状态：在线")
	fmt.Printf("用户信息：%s\n", string(jsonBytes))
	return nil
}
