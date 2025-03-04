package main

import (
	"bytes"
	"testing"
)

func Benchmark_withoutPool(b *testing.B) {
	user := User{"Alice", 30}
	for i := 0; i < b.N; i++ {
		withoutPool(user)
	}
}

func Benchmark_withPool(b *testing.B) {
	user := User{"Alice", 30}
	buf = bytes.NewBuffer(make([]byte, 0, 100))
	for i := 0; i < b.N; i++ {
		withPool(user)
	}
}
