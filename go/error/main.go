package main

import (
	"errors"
	"fmt"
)

func main() {
	err := bazz()
	fmt.Println(err)
}

func bazz() error {
	_, err := foo(false)
	if err != nil {
		return err
	}
	if err := foo2(); err != nil {
		fmt.Println(err)
	}
	fmt.Println(err)

	return nil
}

func foo(isErr bool) (int, error) {
	if isErr {
		return 0, errors.New("foo error")
	}
	return 1, nil
}

func foo2() error {
	return errors.New("foo2 error")
}
