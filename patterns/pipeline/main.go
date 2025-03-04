package main

import "fmt"

func main() {
	values := []int{1, 2, 3, 4, 5}
	mul := func(value int) int {
		return value * value
	}

	for value := range process(generate(values...), mul) {
		fmt.Println(value)
	}
}
