package commons

import (
	"runtime"
	"strconv"
	"time"
)

func IsWindows() bool {
	return runtime.GOOS == "windows"
}

func CurrentMilliSecond() string {
	return strconv.FormatInt(time.Now().UnixMilli(), 10)
}

func CurrentSecond() string {
	return strconv.FormatInt(time.Now().Unix(), 10)
}

func GetOsName() string {
	if IsWindows() {
		return GetOsType() + " NT"
	} else {
		return GetOsType()
	}
}

func GetOsType() string {
	return runtime.GOOS
}
