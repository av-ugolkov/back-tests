package main

import (
	"testing"
)

func Benchmark_addToSlice(b *testing.B) {
	b.Run("addToSlice", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			sl := make([]int, 0, 4)
			sl = append(sl, 1, 2, 3)
			sl = addToSlice(sl)
		}
	})
	b.Run("addToSlicePtr", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			sl := make([]int, 0, 4)
			sl = append(sl, 1, 2, 3)
			addToSlicePtr(&sl)
		}
	})
}
