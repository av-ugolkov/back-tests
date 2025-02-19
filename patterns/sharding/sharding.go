package main

import (
	"crypto/sha1"
	"sync"
)

type ShardType interface {
	string | int | float64
}

type Shard[T ShardType] struct {
	sync.RWMutex
	m map[string]T
}

type ShardedMap[T ShardType] []*Shard[T]

func New[T ShardType](nshards int) ShardedMap[T] {
	shards := make([]*Shard[T], nshards)
	for i := 0; i < nshards; i++ {
		shard := make(map[string]T)
		shards[i] = &Shard[T]{m: shard}
	}

	return shards
}

func (m ShardedMap[T]) Get(key string) any {
	shard := m.getShard(key)
	shard.RLock()
	defer shard.RUnlock()

	return shard.m[key]
}

func (m ShardedMap[T]) Set(key string, value T) {
	shard := m.getShard(key)
	shard.Lock()
	defer shard.Unlock()

	shard.m[key] = value
}

func (m ShardedMap[T]) Delete(key string) {
	shard := m.getShard(key)
	shard.Lock()
	defer shard.Unlock()

	delete(shard.m, key)
}

func (n ShardedMap[T]) Contains(key string) bool {
	shard := n.getShard(key)
	shard.RLock()
	defer shard.RUnlock()

	_, ok := shard.m[key]
	return ok
}

func (m ShardedMap[T]) getShardIndex(key string) int {
	chechsum := sha1.Sum([]byte(key))
	hash := int(chechsum[17]) //random byte for the hash role
	//hash := int(chechsum[13])<<8 | int(chechsum[17])
	return hash % len(m)
}

func (m ShardedMap[T]) Keys() []string {
	keys := make([]string, 0, len(m))

	var mx sync.Mutex
	var wg sync.WaitGroup
	wg.Add(len(m))

	for _, shard := range m {
		go func(s *Shard[T]) {
			s.RLock()

			for key := range s.m {
				mx.Lock()
				keys = append(keys, key)
				mx.Unlock()
			}

			s.RUnlock()
			wg.Done()
		}(shard)
	}

	wg.Wait()

	return keys

}

func (m ShardedMap[T]) getShard(key string) *Shard[T] {
	ind := m.getShardIndex(key)
	return m[ind]
}
