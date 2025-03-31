package main

import (
	"context"
	"log"
	"time"
)

func main() {
	client := NewClient()
	defer client.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := client.SayHello(ctx, "Alex")
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", resp)
}
