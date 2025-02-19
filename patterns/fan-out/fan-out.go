package main

type ChanType interface {
	string | int | float64
}

func Split[T ChanType](source <-chan T, n int) []<-chan T {
	chOut := make([]<-chan T, 0, n)
	for i := 0; i < n; i++ {
		ch := make(chan T)
		chOut = append(chOut, ch)
		go func() {
			defer close(ch)
			for v := range source {
				ch <- v
			}
		}()
	}

	return chOut
}
