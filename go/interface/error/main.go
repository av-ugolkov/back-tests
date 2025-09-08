package main

import "fmt"

type MyError string

func (e MyError) Error() string { return string(e) }

var _ error = (*MyError)(nil) // interface implementation check

func returnMyError() *MyError {
	return nil
}

func returnNil() error {
	return returnMyError()
}

func main() {
	x := returnNil()
	if x == nil {
		fmt.Println("nil!")
	} else {
		fmt.Println(x)
	}
}
