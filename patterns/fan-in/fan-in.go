package main

import "sync"

type ChanType interface {
	string | int | float64
}

func Funnel[T ChanType](chans ...<-chan T) <-chan T {
	chOut := make(chan T)

	var wg sync.WaitGroup
	wg.Add(len(chans))

	for _, c := range chans {
		go func(c <-chan T) {
			defer wg.Done()
			for v := range c {
				chOut <- v
			}
		}(c)
	}

	go func() {
		wg.Wait()
		close(chOut)
	}()

	return chOut
}
