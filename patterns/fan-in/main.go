package main

import (
	"fmt"
	"log/slog"
	"time"
)

func main() {
	sources := make([]<-chan int, 0, 3)

	for i := 0; i < 3; i++ {
		ch := make(chan int)
		sources = append(sources, ch)

		go func() {
			defer close(ch)

			for i := 0; i < 5; i++ {
				ch <- i
				time.Sleep(time.Second)
			}
		}()
	}

	dest := Funnel(sources...)
	for d := range dest {
		slog.Info(fmt.Sprintf("result: %d", d))
	}
}
