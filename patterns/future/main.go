package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	future := SlowFunction[string](ctx)

	res, err := future.Result()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)
}

func SlowFunction[T FutureType](ctx context.Context) Future[T] {
	resCh := make(chan T)
	errCh := make(chan error)
	var zero T
	go func() {
		select {
		case <-time.After(2 * time.Second):
			switch any(zero).(type) {
			case int:
				resCh <- any(11).(T)
				errCh <- nil
			case string:
				resCh <- any("hello").(T)
				errCh <- nil
			case float64:
				resCh <- any(1.1).(T)
				errCh <- nil
			default:
				resCh <- zero
				errCh <- fmt.Errorf("unknown type")
			}
		case <-ctx.Done():
			resCh <- zero
			errCh <- ctx.Err()
		}
	}()

	return &InnerFuture[T]{resCh: resCh, errCh: errCh}
}
