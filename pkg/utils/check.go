package utils

import (
	"errors"
	"reflect"
)

func IsEmpty(value any) bool {
	if value == nil {
		return true
	}
	v := reflect.ValueOf(value)
	switch v.Kind() {
	case reflect.String, reflect.Slice, reflect.Array, reflect.Map, reflect.Chan:
		return v.Len() == 0
	case reflect.Ptr, reflect.Interface:
		return v.IsNil()
	default:
		return false
	}
}

func DefaultIfEmpty[T any](value, defaultValue T) T {
	if IsEmpty(value) {
		return defaultValue
	}
	return value
}

func Require(value any, msg string) error {
	if IsEmpty(value) {
		return errors.New(msg)
	}
	return nil
}
