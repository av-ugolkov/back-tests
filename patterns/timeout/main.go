package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	ctxt, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()
	timeout := Timeout(Slow)
	res, err := timeout(ctxt, "some input")
	fmt.Println(res, err)
}

func Slow(arg string) (string, error) {
	time.Sleep(2 * time.Second)
	return arg, nil
}
