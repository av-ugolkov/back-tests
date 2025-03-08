package main

import "testing"

func Benchmark_changeStrByReflect(t *testing.B) {
	s := new(string)
	*s = "Hello Benchmark"
	for i := 0; i < t.N; i++ {
		changeStrByReflect(s)
	}
}

func Benchmark_changeStrByReflectV(t *testing.B) {
	s := "Hello Benchmark"
	for i := 0; i < t.N; i++ {
		changeStrByReflectV(s)
	}
}

func Benchmark_changeStr(t *testing.B) {
	s := new(string)
	*s = "Hello Benchmark"
	for i := 0; i < t.N; i++ {
		changeStr(s)
	}
}

func Benchmark_changeStrV(t *testing.B) {
	s := "Hello Benchmark"
	for i := 0; i < t.N; i++ {
		changeStrV(s)
	}
}

func Benchmark_unsafeChangeStr(t *testing.B) {
	s := new(string)
	*s = "Hello Benchmark"
	for i := 0; i < t.N; i++ {
		unsafeChangeStr(s)
	}
}

func Benchmark_changeStrByte(t *testing.B) {
	s := new(string)
	*s = "Hello Benchmark"
	for i := 0; i < t.N; i++ {
		changeStrByte(s)
	}
}
