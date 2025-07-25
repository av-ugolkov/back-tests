package main

import (
	"context"
	"log"
	"time"
)

func main() {
	log.Println("work pool")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	data := make(chan int)

	wp := CreatePool(3)
	out := wp.Start(ctx, data)

	go func() {
		for i := 1; i <= 10; i++ {
			select {
			case <-ctx.Done():
				return
			default:
			}
			data <- i
		}
		close(data)
	}()

	for data := range out {
		log.Printf("data: %v", data)
	}
}
