package main

import (
	"context"
	"errors"
	"sync"
	"time"
)

type CircuitType interface {
	string | int | float64
}

type Circuit[T CircuitType] func(context.Context) (T, error)

func Breaker[T CircuitType](circuit Circuit[T], failureThreshold uint) Circuit[T] {
	var consecutiveFailures int = 0
	var lastAttempt = time.Now()
	var m sync.RWMutex
	var zero T

	return func(ctx context.Context) (T, error) {
		m.RLock()
		d := consecutiveFailures - int(failureThreshold)

		if d >= 0 {
			adding := time.Second * 2 << d
			shouldRetryAt := lastAttempt.Add(adding)
			if !time.Now().After(shouldRetryAt) {
				m.RUnlock()
				return zero, errors.New("circuit is open")
			}
		}

		m.RUnlock()
		response, err := circuit(ctx)
		m.Lock()
		defer m.Unlock()

		lastAttempt = time.Now()
		if err != nil {
			consecutiveFailures++
			return zero, err
		}

		consecutiveFailures = 0

		return response, nil
	}
}
