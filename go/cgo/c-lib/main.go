package main

import (
	"fmt"

	"github.com/av-ugolkov/backend-examples/go/cgo/c-lib/sum"
)

func main() {
	fmt.Println("sum c lib")
	ss := sum.Sum(5, 9)
	fmt.Println(ss)
}
