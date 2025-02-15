package utils

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func AnyToError(value any) error {
	if value == nil {
		return nil
	}
	if err, ok := value.(error); ok {
		return err
	}
	return errors.New(fmt.Sprintf("unexpected type: %T", value))
}

func StringToInt(value string) (int, error) {
	return strconv.Atoi(value)
}

func StringToUintSlice(s string) []uint {
	var result []uint
	for _, str := range strings.Split(s, ",") {
		if num, err := strconv.ParseUint(str, 10, 64); err == nil {
			result = append(result, uint(num))
		}
	}
	return result
}
