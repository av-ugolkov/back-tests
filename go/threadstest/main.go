package main

import (
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(1)

	a := []int{}
	i := 0

	for {
		a = append(a, i)
		i++

		if i == 10_000_000_000 {
			a = nil
			i = 0
			break
		}
	}
}
