package main

import "fmt"

func process(closeCh <-chan struct{}) <-chan struct{} {
	closeDoneCh := make(chan struct{})

	go func() {
		defer close(closeDoneCh)

		for {
			select {
			case <-closeCh:
				return
			default:
				fmt.Println("processing...")
			}
		}
	}()

	return closeDoneCh
}
