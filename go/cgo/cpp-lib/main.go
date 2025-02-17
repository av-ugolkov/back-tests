package main

import (
	"fmt"

	"github.com/av-ugolkov/backend-examples/go/cgo/cpp-lib/sum"
)

func main() {
	fmt.Println("sum cpp-lib")
	cppSum := sum.New()
	defer cppSum.Destroy()
	ss := cppSum.Sum(9, 6)
	fmt.Println(ss)
}
