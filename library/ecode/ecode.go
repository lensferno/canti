package ecode

import "fmt"

type ErrCode struct {
	Code int
	Msg  string
}

func NewErrCode(code int, msg string) ErrCode {
	return ErrCode{
		Code: code,
		Msg:  msg,
	}
}

func (e ErrCode) Error() string {
	return fmt.Sprintf("err %d: %s", e.Code, e.Msg)
}
