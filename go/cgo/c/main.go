package main

import (
	"fmt"

	"github.com/av-ugolkov/back-tests/go/cgo/c/sum"
)

func main() {
	fmt.Println("sum c")
	ss := sum.Sum(5, 2)
	fmt.Println(ss)
}
