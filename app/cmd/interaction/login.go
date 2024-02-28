package interaction

import (
	"fmt"
	"github.com/AlecAivazis/survey/v2"
)

type LoginInteraction struct {
	Username  string
	Password  string
	Reconnect bool
	Method    string
	Config    string
}

func NewLoginInteraction() *LoginInteraction {
	return &LoginInteraction{}
}

func (s *LoginInteraction) Ask() error {
	question := []*survey.Question{
		{
			Name:     "username",
			Prompt:   &survey.Input{Message: "用户名？"},
			Validate: survey.Required,
		}, {
			Name:     "password",
			Prompt:   &survey.Password{Message: "密码？"},
			Validate: survey.Required,
		}, {
			Name: "method",
			Prompt: &survey.Select{
				Message: "认证方式？（默认: web，目前仅web方式可用）",
				Options: []string{"web" /*, "pppoe"*/},
				Default: "web",
				VimMode: false,
				Description: func(value string, index int) string {
					switch value {
					case "web":
						return "默认的方式，一般仅限无线网络登录使用，有线网络请使用pppoe"
					case "pppoe":
						return "使用pppoe进行登录认证，仅限有线网络登录使用，无线网络（wifi6和wifi6-edu）请使用web方式"
					default:
						return ""
					}
				},
			},
		}, {
			Name: "reconnect",
			Prompt: &survey.Confirm{
				Message: "如果检测到掉线，是否重连？（默认：Yes）",
				Default: true,
			},
		},
	}

	err := survey.Ask(question, s)
	fmt.Println("----------------")
	if err != nil {
		return err
	}

	return nil
}
