package main

type ChanType interface {
	string | int | float64
}

func Tee[T ChanType](inputCh <-chan T, n int) []<-chan T {
	outputChs := make([]chan T, n)
	for i := 0; i < n; i++ {
		outputChs[i] = make(chan T)
	}

	go func() {
		for value := range inputCh {
			for i := 0; i < n; i++ {
				outputChs[i] <- value
			}
		}

		for _, ch := range outputChs {
			close(ch)
		}
	}()

	resultChs := make([]<-chan T, n)
	for i := 0; i < n; i++ {
		resultChs[i] = outputChs[i]
	}

	return resultChs
}
