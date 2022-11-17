package error

type NotNeedLogin struct{}

func (receiver *NotNeedLogin) Error() string {
	return "Network is available, needn't login."
}

type PasswordWrongError struct{}

func (receiver PasswordWrongError) Error() string {
	return "The password is wrong."
}
