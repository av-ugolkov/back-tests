package main

import (
	"context"
	"time"
)

type RetryType interface {
	string | int | float64
}

type Effector[T RetryType] func(context.Context) (T, error)

func Retry[T RetryType](effector Effector[T], retries int, delay time.Duration) Effector[T] {
	return func(ctx context.Context) (T, error) {
		var zero T
		for r := 0; ; r++ {
			response, err := effector(ctx)
			if err == nil || r >= retries {
				return response, err
			}

			select {
			case <-time.After(delay):
			case <-ctx.Done():
				return zero, ctx.Err()
			}
		}
	}
}
