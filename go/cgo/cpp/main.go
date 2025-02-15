package main

import (
	"fmt"

	"github.com/av-ugolkov/backend-examples/go/cgo/cpp/sum"
)

func main() {
	fmt.Println("sum cpp")
	cppSum := sum.New()
	defer cppSum.Destroy()
	ss := cppSum.Sum(5, 9)
	fmt.Println(ss)
}
