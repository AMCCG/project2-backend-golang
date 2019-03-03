package utils

import (
	"reflect"
	"strings"
)

func Reflection(entity interface{}) string {
	if t := reflect.TypeOf(entity); t.Kind() == reflect.Ptr {
		return strings.ToLower(t.Elem().Name())
	} else {
		return strings.ToLower(t.Name())
	}
}
