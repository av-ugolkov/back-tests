package main

import (
	"fmt"
	"sync"
)

func main() {
	source := make(chan int)
	dests := Split(source, 5)

	go func() {
		for i := 0; i < 1000; i++ {
			source <- i
		}
		close(source)
	}()

	var wg sync.WaitGroup
	wg.Add(len(dests))
	for i, ch := range dests {
		go func(i int, d <-chan int) {
			defer wg.Done()

			for val := range d {
				fmt.Printf("%d: %d\n", i, val)
			}
		}(i, ch)
	}

	wg.Wait()
}
