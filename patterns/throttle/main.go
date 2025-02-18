package main

import (
	"context"
	"fmt"
	"log/slog"
	"time"
)

func main() {
	fnEffector := func(ctx context.Context) (int, error) {
		return 1, nil
	}

	ctx := context.Background()
	fn := Throttle(fnEffector, 5, 1, 100*time.Millisecond)
	for i := 0; i < 7; i++ {
		data, err := fn(ctx)
		checkError(data, err)
	}
	time.Sleep(1 * time.Second)
	data, err := fn(ctx)
	checkError(data, err)
}

func checkError(data any, err error) {
	if err != nil {
		slog.Error(err.Error())
	} else {
		slog.Info(fmt.Sprintf("result: %v", data))
	}
}
