package main

import "testing"

var ret1 int
var ret2 int

func BenchmarkPart1(b *testing.B) {
	r := 0
	for n := 0; n < b.N; n++ {
		r = part1()
	}
	ret1 = r
}

func BenchmarkPart2(b *testing.B) {
	r := 0
	for n := 0; n < b.N; n++ {
		r = part2()
	}
	ret2 = r
}
