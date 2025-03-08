package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	s := "Hello Go"
	changeStr(&s)
	fmt.Println(s)
	fmt.Println(changeStrV(s))
	changeStrByReflect(&s)
	fmt.Println(s)
	fmt.Println(changeStrByReflectV(s))
	unsafeChangeStr(&s)
	fmt.Println(s)
	changeStrByte(&s)
	fmt.Println(s)
}

func changeStrByReflect(s *string) {
	ref := reflect.ValueOf(s).Elem()
	ref.Set(reflect.ValueOf("Hello Reflect"))
}

func changeStrByReflectV(s string) string {
	ref := reflect.ValueOf(&s).Elem()
	ref.Set(reflect.ValueOf("Hello Reflect Value"))
	return ref.String()
}

func changeStr(s *string) {
	*s = "Hello New String"
}

func changeStrV(s string) string {
	s = "Hello New String Value"
	return s
}

func unsafeChangeStr(s *string) {
	*s = string([]byte(*s))
	strHeader := (*struct {
		data unsafe.Pointer
		len  int
	})(unsafe.Pointer(s))
	strBytes := unsafe.Slice((*byte)(strHeader.data), strHeader.len)
	strBytes[0] = 'T'
}

func changeStrByte(s *string) {
	b := []byte(*s)
	b[0] = 'T'
}
