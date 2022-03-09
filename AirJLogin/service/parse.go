package service

import (
	"strconv"
	"strings"
)

// UnicodeStrToEntity 用于转义json中的中文unicode编码
func UnicodeStrToEntity(s string) string {
	s1, _ := strconv.Unquote(strings.Replace(strconv.Quote(s), `\\u`, `\u`, -1))
	return s1
}
