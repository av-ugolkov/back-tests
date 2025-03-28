package main

import (
	"fmt"
	"log/slog"
	"sync"
	"time"
)

const defaultValue = 42

func main() {
	m := make(map[int]int, 1)
	m[defaultValue] = defaultValue

	var rm sync.RWMutex
	change := false

	for i := 0; i < 5; i++ {
		go func() {
			run := true
			go func() {
				time.Sleep(300 * time.Millisecond)
				run = false
			}()
			for run {
				rm.RLock()
				if v := m[defaultValue]; v != defaultValue || change {
					slog.Info(fmt.Sprintf("%v", v))
					if v != defaultValue {
						run = false
					}
				}
				rm.RUnlock()
			}
		}()
	}
	time.Sleep(50 * time.Millisecond)

	go func() {
		slog.Info("start change")
		change = true
		rm.Lock()
		m[defaultValue] = 24
		rm.Unlock()
	}()

	time.Sleep(400 * time.Millisecond)
}
