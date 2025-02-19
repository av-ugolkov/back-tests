/*
Also known as Promises or Delays
*/
package main

import (
	"sync"
)

type FutureType interface {
	string | int | float64
}

type Future[T FutureType] interface {
	Result() (T, error)
}

type InnerFuture[T FutureType] struct {
	once sync.Once
	wg   sync.WaitGroup

	res   T
	err   error
	resCh <-chan T
	errCh <-chan error
}

func (f *InnerFuture[T]) Result() (T, error) {
	f.once.Do(func() {
		f.wg.Add(1)
		defer f.wg.Done()
		f.res, f.err = <-f.resCh, <-f.errCh
	})

	f.wg.Wait()

	return f.res, f.err
}
