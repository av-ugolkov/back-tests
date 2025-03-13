package main

import (
	"fmt"
	"time"
)

func main() {
	go fnExec()

	time.Sleep(1 * time.Second)

	fmt.Println("main")
}

func fnExec() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovered from", r)
		}
	}()
	fnPanic()

	fmt.Println("fnExec")
}

func fnPanic() {
	panic("panic")
}
