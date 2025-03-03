package main

import "fmt"

func main() {
	ch := make(chan int)

	go func() {
		defer close(ch)
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()

	isOdd := func(value int) bool {
		return value%2 != 0
	}

	for num := range Filter(ch, isOdd) {
		fmt.Println(num)
	}
}
