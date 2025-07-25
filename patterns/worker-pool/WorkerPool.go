package main

import (
	"context"
	"log"
	"math/rand"
	"sync"
	"time"
)

type WorkerPool struct {
	countJobs int
}

func CreatePool(count int) *WorkerPool {
	return &WorkerPool{
		countJobs: count,
	}
}

func (p *WorkerPool) Start(ctx context.Context, in chan int) chan int {
	out := make(chan int, p.countJobs)

	var wg sync.WaitGroup
	for i := 0; i < p.countJobs; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			worker(ctx, i+1, in, out)
		}()
	}

	go func() {
		wg.Wait()
		defer close(out)
	}()

	return out
}

func worker(ctx context.Context, id int, in, out chan int) {
	for {
		select {
		case <-ctx.Done():
			log.Printf("jod [%v] stop", id)
			return
		case value, ok := <-in:
			if !ok {
				return
			}
			wait := 100 + rand.Int31n(999)
			select {
			case <-ctx.Done():
				log.Printf("jod [%v] cancelled before processing value [%v]", id, value)
				return
			case <-time.After(time.Duration(wait) * time.Millisecond):
				log.Printf("job [%v] with value [%v] done", id, value)
				out <- value
			}
		}
	}
}
