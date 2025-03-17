package main

import "sync"

func main() {
	sl := make([]int, 2)
	m := make(map[int]int, 2)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		sl[0] = 1
		m[0] = 1
	}()

	go func() {
		defer wg.Done()
		sl[1] = 2
		m[1] = 2
	}()

	wg.Wait()
}
