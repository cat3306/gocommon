package strutil

import (
	"fmt"
	"reflect"
	"unsafe"
)

//zero copy
func BytesToString(raw []byte) string {
	return *(*string)(unsafe.Pointer(&raw))
}
func StringToBytes(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(&struct {
		string
		Cap int
	}{s, len(s)}))
}

func StructToStringArray(o interface{}) []string {
	column := make([]string, 0)
	v := reflect.ValueOf(o)
	for i := 0; i < v.NumField(); i++ {
		column = append(column, fmt.Sprintf("%v", v.Field(i).Interface()))
	}
	return column
}
