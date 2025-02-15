package main

import (
	"fmt"

	"github.com/av-ugolkov/backend-examples/go/cgo/c/sum"
)

func main() {
	fmt.Println("sum c")
	ss := sum.Sum(5, 6)
	fmt.Println(ss)
}
