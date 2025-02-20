package main

import "sync"

func Chord(sources ...<-chan int) <-chan []int {
	type input struct {
		idx, input int
	}

	dest := make(chan []int)

	inputs := make(chan input)

	wg := sync.WaitGroup{}
	wg.Add(len(sources))

	for i, ch := range sources {
		go func(i int, ch <-chan int) {
			defer wg.Done()

			for n := range ch {
				inputs <- input{i, n}
			}
		}(i, ch)
	}

	go func() {
		wg.Wait()
		close(inputs)
	}()

	go func() {
		res := make([]int, len(sources))
		sent := make([]bool, len(sources))
		count := len(sources)

		for r := range inputs {
			res[r.idx] = r.input

			if !sent[r.idx] {
				sent[r.idx] = true
				count--
			}

			if count == 0 {
				c := make([]int, len(res))
				copy(c, res)
				dest <- c

				count = len(sources)
				clear(sent)
			}
		}

		close(dest)
	}()

	return dest
}
