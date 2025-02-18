package main

import (
	"context"
	"testing"
)

func TestMain(t *testing.T) {
	work(context.Background(), "alice")
}
