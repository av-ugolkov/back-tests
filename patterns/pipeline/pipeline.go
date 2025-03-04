package main

type ChanType interface {
	string | int | float64
}

func process[T ChanType](input <-chan T, action func(T) T) <-chan T {
	output := make(chan T)
	go func() {
		defer close(output)
		for value := range input {
			output <- action(value)
		}
	}()

	return output
}

func generate[T ChanType](values ...T) <-chan T {
	output := make(chan T)

	go func() {
		defer close(output)
		for _, value := range values {
			output <- value
		}
	}()

	return output
}
