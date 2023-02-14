package utils

import (
	"strconv"
	"strings"
	"time"
)

// 把字符串时间段，转换为duration对象
func ParseDuration(d string) (time.Duration, error) {
	d = strings.TrimSpace(d)
	// 只能解析最大到小时
	dr, err := time.ParseDuration(d)
	if err == nil {
		return dr, nil
	}
	// 包含d，则需要手动解析
	if strings.Contains(d, "d") {
		index := strings.Index(d, "d")
		day, _ := strconv.Atoi(d[:index])
		dr = time.Hour * 24 * time.Duration(day)
		ndr, err := time.ParseDuration(d[index+1:])
		if err != nil {
			return dr, nil
		}
		return dr + ndr, nil
	}
	// 不包含d，还解析不了,直接做数字处理
	dv, err := strconv.ParseInt(d, 10, 64)
	return time.Duration(dv), err
}
