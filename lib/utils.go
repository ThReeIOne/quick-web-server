package lib

import (
	_ "image/jpeg"
	_ "image/png"
	"quick_web_golang/config"
	"strconv"
	"strings"
	"time"
)

func IsDev() bool {
	debug, _ := strconv.ParseBool(config.Get(config.IsDev))
	return debug
}

func IsEnableNetwork() bool {
	enable, _ := strconv.ParseBool(config.Get(config.EnableNetwork))
	return enable
}

func ParseDuration(d string) (time.Duration, error) {
	d = strings.TrimSpace(d)
	dr, err := time.ParseDuration(d)
	if err == nil {
		return dr, nil
	}
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

	dv, err := strconv.ParseInt(d, 10, 64)
	return time.Duration(dv), err
}

func RandomId() string {
	id := time.Now().Format("20060102150405.000000")
	return strings.ReplaceAll(id, ".", "")
}

func IntToBool(i int) bool {
	return i > 0
}

func BoolToInt(boo bool) int {
	if boo {
		return 1
	}
	return 0
}

func StrToPtr(s string) *string {
	return &s
}
