package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	<-or(
		time.After(2*time.Second),
		time.After(1*time.Second),
		time.After(900*time.Millisecond),
		time.After(550*time.Millisecond),
		time.After(650*time.Millisecond),
		time.After(750*time.Millisecond),
		time.After(850*time.Millisecond),
	)

	fmt.Printf("called after: %s\n", time.Since(start))
}
