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

	checkErr := func(ind int, d any, err error) {
		if err != nil {
			slog.Error(fmt.Sprintf("%d error: %v", ind, err))
		} else {
			slog.Info(fmt.Sprintf("%d result: %v", ind, d))
		}
	}

	fn := DebounceFirst(circuit, time.Second*2)

	v, err := fn(context.WithValue(context.Background(), "key", 1))
	checkErr(1, v, err)
	time.Sleep(1 * time.Second)
	v, err = fn(context.WithValue(context.Background(), "key", 2))
	checkErr(1, v, err)
	time.Sleep(1 * time.Second)
	v, err = fn(context.WithValue(context.Background(), "key", 3))
	checkErr(1, v, err)
	time.Sleep(1 * time.Second)
	v, err = fn(context.WithValue(context.Background(), "key", 4))
	checkErr(1, v, err)
}
