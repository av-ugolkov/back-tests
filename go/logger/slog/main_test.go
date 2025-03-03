package main

import "testing"

func BenchmarkSlog(b *testing.B) {
	for i := 0; i < b.N; i++ {
		main()
	}
}
