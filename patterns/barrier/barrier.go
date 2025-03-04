package main

import "sync"

type Barrier struct {
	count int
	size  int
	mutex sync.Mutex

	beforeCh chan struct{}
	afterCh  chan struct{}
}

func NewBarrier(size int) Barrier {
	return Barrier{
		size:     size,
		beforeCh: make(chan struct{}, size),
		afterCh:  make(chan struct{}, size),
	}
}

func (b *Barrier) Before() {
	b.mutex.Lock()

	b.count++
	if b.count == b.size {
		for i := 0; i < b.size; i++ {
			b.beforeCh <- struct{}{}
		}
	}

	b.mutex.Unlock()
	<-b.beforeCh
}

func (b *Barrier) After() {
	b.mutex.Lock()

	b.count--
	if b.count == 0 {
		for i := 0; i < b.size; i++ {
			b.afterCh <- struct{}{}
		}
	}

	b.mutex.Unlock()
	<-b.afterCh
}
