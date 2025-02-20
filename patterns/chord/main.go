package main

import (
	"fmt"
	"log/slog"
	"time"
)

func main() {
	expected := [][]int{
		{2, 1, 1},
		{4, 2, 2},
		{6, 3, 3},
		{8, 4, 4},
		{10, 5, 5},
	}

	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	go func() {
		for n := 1; n <= 5; n++ {
			ch1 <- n
			ch1 <- n * 2
			ch2 <- n
			ch3 <- n
			time.Sleep(time.Second)
		}

		close(ch1)
		close(ch2)
		close(ch3)
	}()

	res := Chord(ch1, ch2, ch3)
	idx := 0

	for ss := range res {
		fmt.Println(ss)

		for i, s := range ss {
			if expected[idx][i] != s {
				slog.Error(fmt.Sprintf("Expected: %v; Got: %v", expected[idx], ss))
				return
			}
		}

		idx++
	}
}
