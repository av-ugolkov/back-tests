package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var i8 int8
	var i16 int16
	var i32 int32
	var i64 int64
	fmt.Printf("int8: %v\nint16: %v\nint32: %v\nint64: %v\n", unsafe.Sizeof(i8), unsafe.Sizeof(i16), unsafe.Sizeof(i32), unsafe.Sizeof(i64))

}
