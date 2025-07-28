package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type ThrottleType interface {
	string | int | float64
}

type Effector[T ThrottleType] func(context.Context) (T, error)

func Throttle[T ThrottleType](e Effector[T], max uint, refill uint, d time.Duration) Effector[T] {
	var tokens = max
	var once sync.Once
	var zero T

	return func(ctx context.Context) (T, error) {
		if ctx.Err() != nil {
			return zero, ctx.Err()
		}
		once.Do(func() {
			ticker := time.NewTicker(d)
			go func() {
				defer ticker.Stop()

				for {
					select {
					case <-ctx.Done():
						return

					case <-ticker.C:
						tokens = tokens + refill
						if tokens > max {
							tokens = max
						}
					}
				}
			}()
		})
		if tokens <= 0 {
			return zero, fmt.Errorf("too many calls")
		}

		tokens--
		return e(ctx)
	}
}
