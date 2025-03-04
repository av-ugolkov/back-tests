package main

import (
	"fmt"
	"time"
)

func main() {
	closeCh := make(chan struct{})
	closeDoneCh := process(closeCh)

	time.Sleep(50 * time.Microsecond)
	close(closeCh)
	<-closeDoneCh

	fmt.Println("terminated")
}
