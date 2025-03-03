package main

import (
	"log/slog"
)

func main() {
	for i := 1; i <= 10; i++ {
		slog.Info("Testing sampling", "index", i)
	}
}
