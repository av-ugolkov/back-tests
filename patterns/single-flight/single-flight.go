package main

import "sync"

type call struct {
	err error
	val any

	done chan struct{}
}

type Singleflight struct {
	mutex sync.Mutex
	calls map[string]*call
}

func NewSingleFlight() *Singleflight {
	return &Singleflight{
		calls: make(map[string]*call),
	}
}

func (s *Singleflight) Do(key string, action func() (any, error)) (any, error) {
	s.mutex.Lock()
	if call, found := s.calls[key]; found {
		s.mutex.Unlock()
		return s.wait(call)
	}

	call := &call{
		done: make(chan struct{}),
	}

	s.calls[key] = call
	s.mutex.Unlock()

	go func() {
		defer func() {
			s.mutex.Lock()
			close(call.done)
			delete(s.calls, key)
			s.mutex.Unlock()
		}()

		call.val, call.err = action()
	}()

	return s.wait(call)
}

func (s *Singleflight) wait(call *call) (any, error) {
	<-call.done
	return call.val, call.err
}
