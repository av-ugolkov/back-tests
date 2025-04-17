package main

import "testing"

func Benchmark_split(t *testing.B) {
	str := "qwer;asdf;zxcv;wert;sdfg;xcvb;erty;dfgh;cvbn"
	sep := ";"
	for i := 0; i < t.N; i++ {
		split(str, sep)
	}
}

func Benchmark_cut(t *testing.B) {
	str := "qwer;asdf;zxcv;wert;sdfg;xcvb;erty;dfgh;cvbn"
	sep := ";"
	for i := 0; i < t.N; i++ {
		cut(str, sep)
	}
}

func Benchmark_index(t *testing.B) {
	str := "qwer;asdf;zxcv;wert;sdfg;xcvb;erty;dfgh;cvbn"
	sep := ";"
	for i := 0; i < t.N; i++ {
		index(str, sep)
	}
}
