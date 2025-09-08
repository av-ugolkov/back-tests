package main

import "fmt"

type MyInt int

func main() {
	var v any = 11
	mi := v.(MyInt) //panic: interface conversion: interface {} is int, not main.MyInt
	fmt.Println(mi)
}
