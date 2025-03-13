package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	go loop(ctx, "go")
	loop(ctx, "main")
}

func loop(ctx context.Context, v string) {
	for {
		fmt.Println("Working...", v)

		select {
		case <-ctx.Done():
			fmt.Println("Done", v)
			return
		default:
		}

		runtime.Gosched()
	}
}
