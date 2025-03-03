package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		waitNanoseconds()
		waitMicroseconds()
		waitMiliseconds()
		waitSeconds()
	}
}

func waitNanoseconds() {
	var wg sync.WaitGroup
	fmt.Print("waitNanoseconds ")
	for _, n := range []int{3, 1, 2} {
		wg.Add(1)
		go func() {
			defer wg.Done()
			time.Sleep(time.Duration(n) * time.Nanosecond)
			fmt.Printf("%d ", n)
		}()
	}
	wg.Wait()
	fmt.Println()
}

func waitMicroseconds() {
	var wg sync.WaitGroup
	fmt.Print("waitMicroseconds ")
	for _, n := range []int{3, 1, 2} {
		wg.Add(1)
		go func() {
			defer wg.Done()
			time.Sleep(time.Duration(n) * time.Microsecond)
			fmt.Printf("%d ", n)
		}()
	}
	wg.Wait()
	fmt.Println()
}

func waitMiliseconds() {
	var wg sync.WaitGroup
	fmt.Print("waitMiliseconds ")
	for _, n := range []int{3, 1, 2} {
		wg.Add(1)
		go func() {
			defer wg.Done()
			time.Sleep(time.Duration(n) * time.Millisecond)
			fmt.Printf("%d ", n)
		}()
	}
	wg.Wait()
	fmt.Println()
}

func waitSeconds() {
	var wg sync.WaitGroup
	fmt.Print("waitSeconds ")
	for _, n := range []int{3, 1, 2} {
		wg.Add(1)
		go func() {
			defer wg.Done()
			time.Sleep(time.Duration(n) * time.Second)
			fmt.Printf("%d ", n)
		}()
	}
	wg.Wait()
	fmt.Println()
}
