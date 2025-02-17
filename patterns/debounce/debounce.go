package main

import (
	"context"
	"sync"
	"time"
)

type CircuitType interface {
	string | int | float64
}

// Circuit is a function that returns a value and an error
type Circuit[T CircuitType] func(context.Context) (T, error)

/*
Decorator for a circuit that returns the first successful value from a sequence in a given duration
*/
func DebounceFirst[T CircuitType](circuit Circuit[T], d time.Duration) Circuit[T] {
	var threshold time.Time
	var result T
	var err error
	var m sync.Mutex

	return func(ctx context.Context) (T, error) {
		m.Lock()

		if time.Now().Before(threshold) {
			m.Unlock()
			return result, err
		}

		defer func() {
			threshold = time.Now().Add(d)
			m.Unlock()
		}()

		result, err = circuit(ctx)
		if err != nil {
			return result, err
		}

		return result, nil
	}
}
