package main

import "testing"

func BenchmarkZap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		main()
	}
}
