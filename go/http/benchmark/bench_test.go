package main

import (
	"net/http"
	"testing"
)

func BenchmarkServe(b *testing.B) {
	for _, f := range frameworks() {
		for _, c := range benchCases() {
			h := start(f)

			code, err := h.do(c)
			if err != nil || code != http.StatusOK {
				h.stop()
				b.Fatalf("%s/%s warm-up: code=%d err=%v", f.name, c.name, code, err)
			}

			b.Run(f.name+"/"+c.name+"/serial", func(b *testing.B) {
				b.ReportAllocs()
				for b.Loop() {
					code, err := h.do(c)
					if err != nil {
						b.Fatalf("request: %v", err)
					}
					if code != http.StatusOK {
						b.Fatalf("expected 200, got %d", code)
					}
				}
			})

			b.Run(f.name+"/"+c.name+"/parallel", func(b *testing.B) {
				b.ReportAllocs()
				b.RunParallel(func(pb *testing.PB) {
					for pb.Next() {
						code, err := h.do(c)
						if err != nil {
							b.Fatalf("request: %v", err)
						}
						if code != http.StatusOK {
							b.Fatalf("expected 200, got %d", code)
						}
					}
				})
			})

			h.stop()
		}
	}
}
