package commons

import (
	"fmt"
	"regexp"
)

var jqueryJsonRegexp = regexp.MustCompile("\\d+\\((?P<json>.*?)\\)$")

func FilterJQueryPrefix(source string) string {
	result := jqueryJsonRegexp.FindStringSubmatch(source)

	if len(result) < 2 {
		fmt.Errorf("jquery字段提纯失败，原字符串：", source)
		return ""
	} else {
		return result[1]
	}
}
