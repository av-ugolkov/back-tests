package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"time"
)

func main() {
	circuit := func(ctx context.Context) (int, error) {
		v := ctx.Value("key")
		if v.(int) == -1 {
			return -1, errors.New("some error")
		}
		return v.(int), nil
	}

	fn := Breaker(circuit, 2)

	checkErr := func(ind int, d any, err error) {
		if err != nil {
			slog.Error(fmt.Sprintf("%d error: %v", ind, err))
		} else {
			slog.Info(fmt.Sprintf("%d result: %v", ind, d))
		}
	}

	ctx := context.Background()
	for i := 0; i < 3; i++ {
		d, err := fn(context.WithValue(ctx, "key", -1))
		checkErr(1, d, err)
	}
	time.Sleep(3 * time.Second)
	d, err := fn(context.WithValue(ctx, "key", -1))
	checkErr(2, d, err)
	time.Sleep(4 * time.Second)
	d, err = fn(context.WithValue(ctx, "key", 1))
	checkErr(3, d, err)
	d, err = fn(context.WithValue(ctx, "key", -1))
	checkErr(4, d, err)

}
