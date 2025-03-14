package main

import "sync"

type PubSub struct {
	mu       sync.Mutex
	channels map[string][]chan string
}

func NewPubSub() *PubSub {
	return &PubSub{
		channels: make(map[string][]chan string),
	}
}

func (ps *PubSub) Subscribe(topic string) <-chan string {
	ch := make(chan string)
	ps.mu.Lock()
	ps.channels[topic] = append(ps.channels[topic], ch)
	ps.mu.Unlock()
	return ch
}

func (ps *PubSub) Publish(topic, msg string) {
	ps.mu.Lock()
	for _, ch := range ps.channels[topic] {
		ch <- msg
	}
	ps.mu.Unlock()
}

func (ps *PubSub) Close(topic string) {
	ps.mu.Lock()
	for _, ch := range ps.channels[topic] {
		close(ch)
	}
	ps.mu.Unlock()
}
