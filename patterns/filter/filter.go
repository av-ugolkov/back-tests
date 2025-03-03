package main

type ChanType interface {
	string | int | float64
}

func Filter[T ChanType](in <-chan T, predicate func(T) bool) <-chan T {
	outputCh := make(chan T)

	go func() {
		defer close(outputCh)
		for value := range in {
			if predicate(value) {
				outputCh <- value
			}
		}
	}()

	return outputCh
}
