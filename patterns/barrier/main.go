package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(3)

	bootstrap := func() {
		fmt.Println("bootstrap")
	}

	work := func() {
		fmt.Println("work")
	}

	count := 3
	barrier := NewBarrier(count)
	for i := 0; i < count; i++ {
		go func() {
			defer wg.Done()

			for j := 0; j < count; j++ {
				barrier.Before()
				bootstrap()
				barrier.After()
				work()
			}
		}()
	}

	wg.Wait()
}
