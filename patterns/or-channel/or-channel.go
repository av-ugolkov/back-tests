package main

import (
	"time"
)

type ChanType interface {
	string | int | float64 | time.Time
}

func or[T ChanType](channels ...<-chan T) <-chan T {
	switch len(channels) {
	case 0:
		return nil
	case 1:
		return channels[0]
	}

	doneCh := make(chan T)

	go func() {
		defer close(doneCh)
		switch len(channels) {
		case 2:
			select {
			case <-channels[0]:
			case <-channels[1]:
			}
		default:
			select {
			case <-channels[0]:
			case <-channels[1]:
			case <-channels[2]:
			case <-or(channels[3:]...):
			}
		}
	}()

	return doneCh
}
