package core

import (
	"errors"
	"log"
)

var ErrorNoSuchKey = errors.New("no such key")

type KeyValueStore struct {
	m        map[string]string
	transact TransactionLogger
}

func NewKeyValueStore(tl TransactionLogger) *KeyValueStore {
	return &KeyValueStore{
		m:        make(map[string]string),
		transact: tl,
	}
}

func (s *KeyValueStore) Get(key string) (string, error) {
	value, ok := s.m[key]

	if !ok {
		return "", ErrorNoSuchKey
	}

	return value, nil
}

func (s *KeyValueStore) Put(key string, value string) error {
	s.m[key] = value
	s.transact.WritePut(key, value)

	return nil
}

func (s *KeyValueStore) Delete(key string) error {
	delete(s.m, key)
	s.transact.WriteDelete(key)

	return nil
}

func (store *KeyValueStore) Restore() error {
	var err error

	events, errors := store.transact.ReadEvents()
	count, ok, e := 0, true, Event{}

	for ok && err == nil {
		select {
		case err, ok = <-errors:

		case e, ok = <-events:
			switch e.EventType {
			case EventDelete: // Got a DELETE event!
				err = store.Delete(e.Key)
				count++
			case EventPut: // Got a PUT event!
				err = store.Put(e.Key, e.Value)
				count++
			}
		}
	}

	log.Printf("%d events replayed\n", count)

	store.transact.Run()

	go func() {
		for err := range store.transact.Err() {
			log.Print(err)
		}
	}()

	return err
}

type ZeroTransactionLogger struct{}

func (z ZeroTransactionLogger) WriteDelete(key string)                   {}
func (z ZeroTransactionLogger) WritePut(key, value string)               {}
func (z ZeroTransactionLogger) Err() <-chan error                        { return nil }
func (z ZeroTransactionLogger) LastSequence() uint64                     { return 0 }
func (z ZeroTransactionLogger) Run()                                     {}
func (z ZeroTransactionLogger) Wait()                                    {}
func (z ZeroTransactionLogger) Close() error                             { return nil }
func (z ZeroTransactionLogger) ReadEvents() (<-chan Event, <-chan error) { return nil, nil }
