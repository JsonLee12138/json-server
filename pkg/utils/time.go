package utils

import (
	"strconv"
	"strings"
	"time"
)

func ConvertFormat(jsFormat string) string {
	replacer := strings.NewReplacer(
		"YYYY", "2006",
		"YY", "06",
		"MM", "01",
		"M", "1",
		"DD", "02",
		"D", "2",
		"HH", "15",
		"H", "3",
		"hh", "03",
		"h", "3",
		"mm", "04",
		"m", "4",
		"ss", "05",
		"s", "5",
		"A", "PM",
		"a", "pm",
	)
	return replacer.Replace(jsFormat)
}

func ParseDuration(value string) (time.Duration, error) {
	value = strings.TrimSpace(value)
	dr, err := time.ParseDuration(value)
	if err == nil {
		return dr, nil
	}
	index := strings.Index(value, "d")
	if index > -1 {
		d, _ := strconv.Atoi(value[:index])
		dr = time.Duration(d) * 24 * time.Hour
		ndr, err := time.ParseDuration(value[index+1:])
		if err != nil {
			return dr, err
		}
		return dr + ndr, nil
	}
	dv, err := strconv.ParseInt(value, 10, 64)
	return time.Duration(dv), err
}
