package main

import (
	"context"
)

type TimeoutType interface {
	string | int | float64
}

type SlowFunction[T TimeoutType] func(T) (T, error)

type WithContext[T TimeoutType] func(context.Context, T) (T, error)

func Timeout[T TimeoutType](f SlowFunction[T]) WithContext[T] {
	var zero T
	return func(ctx context.Context, arg T) (T, error) {
		chres := make(chan T)
		cherr := make(chan error)
		defer func() {
			close(chres)
			close(cherr)
		}()
		go func() {
			res, err := f(arg)
			if ctx.Err() != nil {
				return
			}
			chres <- res
			cherr <- err
		}()

		select {
		case res := <-chres:
			return res, <-cherr
		case <-ctx.Done():
			return zero, ctx.Err()
		}
	}
}
