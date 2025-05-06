package main

import (
	"testing"
)

func Benchmark_addToSlice(b *testing.B) {
	b.Run("addToSlice", func(b *testing.B) {
		sl := make([]int, 0, 4)
		sl = append(sl, 1, 2, 3)
		for i := 0; i < b.N; i++ {
			sl = addToSlice(sl, i)
		}
	})
	b.Run("addToSlicePtr", func(b *testing.B) {
		sl := make([]int, 0, 4)
		sl = append(sl, 1, 2, 3)
		for i := 0; i < b.N; i++ {
			addToSlicePtr(&sl, i)
		}
	})
}
