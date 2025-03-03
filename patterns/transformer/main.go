package main

import "fmt"

func main() {
	ch := make(chan int)

	go func() {
		defer close(ch)
		for i := 0; i < 5; i++ {
			ch <- i
		}
	}()

	mul := func(value int) int {
		return value * value
	}

	for num := range Transformer(ch, mul) {
		fmt.Println(num)
	}
}
