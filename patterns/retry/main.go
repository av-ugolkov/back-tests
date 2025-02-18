package main

import (
	"context"
	"fmt"
	"log/slog"
	"time"
)

func main() {
	retry := 0
	fnRetry := func(ctx context.Context) (int, error) {
		defer func() {
			retry++
		}()
		if retry < 2 {
			return 0, fmt.Errorf("some error")
		}
		return retry, nil
	}

	fn := Retry(fnRetry, 3, time.Second)
	slog.Info("start retry")
	data, err := fn(context.Background())
	if err != nil {
		slog.Error(err.Error())
	} else {
		slog.Info(fmt.Sprintf("result: %v", data))
	}
}
