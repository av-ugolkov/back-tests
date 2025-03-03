package main

type ChanType interface {
	string | int | float64
}

func Transformer[T ChanType](in <-chan T, action func(T) T) <-chan T {
	outputCh := make(chan T)

	go func() {
		defer close(outputCh)
		for value := range in {
			outputCh <- action(value)
		}
	}()

	return outputCh
}
