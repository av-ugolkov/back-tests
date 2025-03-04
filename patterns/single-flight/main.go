package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	const inFlightRequests = 5
	var wg sync.WaitGroup
	wg.Add(inFlightRequests)

	singleFlight := NewSingleFlight()
	const key = "some_key"
	for i := 0; i < inFlightRequests; i++ {
		go func() {
			defer wg.Done()
			value, err := singleFlight.Do(key, func() (any, error) {
				fmt.Println("single flight")
				time.Sleep(3 * time.Second)
				return "result", nil
			})

			fmt.Println(i, "=", value, err)
		}()
	}

	wg.Wait()
}
