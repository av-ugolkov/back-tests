package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int)

	go func() {
		defer close(ch)
		for i := 0; i < 5; i++ {
			ch <- i
		}
	}()

	var wg sync.WaitGroup
	wg.Add(2)

	channels := Tee(ch, 2)
	go func() {
		defer wg.Done()
		for value := range channels[0] {
			fmt.Println("ch1: ", value)
		}
	}()
	go func() {
		defer wg.Done()
		for value := range channels[1] {
			fmt.Println("ch2: ", value)
		}
	}()

	wg.Wait()
}
