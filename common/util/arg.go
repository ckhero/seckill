package util

import (
	"os"
	"regexp"
	"seckill/common/constant"
	"strings"
)

func GetArg(name, defaultVal string) string {
	args := os.Args;
	pattern := name + constant.SepEqual
	for _, arg := range args{
		if match, _:= regexp.MatchString(pattern, arg); match {
				value := strings.Split(arg, constant.SepEqual)
				if (value[0] == name) {
					return strings.TrimSpace(value[1])
				}
		}
	}
	return strings.TrimSpace(defaultVal)
}