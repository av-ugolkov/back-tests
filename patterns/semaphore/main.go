package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	s := NewSemaphore(10)
	wg.Add(12)
	for i := 0; i < 12; i++ {
		s.Acquire()
		go func() {
			defer func() {
				wg.Done()
				s.Release()
			}()

			fmt.Println("working...")
			time.Sleep(2 * time.Second)
			fmt.Println("exiting...")
		}()
	}

	wg.Wait()
}
