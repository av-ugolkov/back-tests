package main

import "fmt"

func main() {
	fnExec()

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
