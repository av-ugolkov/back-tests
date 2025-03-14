package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ps := NewPubSub()

	subscriber1 := ps.Subscribe("news")
	subscriber2 := ps.Subscribe("news")

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		for msg := range subscriber1 {
			fmt.Println("Subscriber 1 received:", msg)
		}
	}()

	go func() {
		defer wg.Done()
		for msg := range subscriber2 {
			fmt.Println("Subscriber 2 received:", msg)
		}
	}()

	ps.Publish("news", "Breaking News!")
	ps.Publish("news", "Another News!")

	time.Sleep(time.Second)
	ps.Close("news")
	wg.Wait()
}
