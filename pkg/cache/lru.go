package cache

import (
	"container/list"
	"iter"
)

type LRUCache[K comparable, V any] interface {
	Get(key K) (value V, ok bool)
	Put(key K, value V)
	Size() int
	All() iter.Seq2[K, V]
	Clear()
}

var _ LRUCache[any, any] = (*lrucache[any, any])(nil)

type node[K comparable, V any] struct {
	key   K
	value V
}

type lrucache[K comparable, V any] struct {
	keyToElement map[K]*list.Element
	linkedList   *list.List
	capacity     int
}

func NewLRUCache[K comparable, V any](capacity int) LRUCache[K, V] {
	return &lrucache[K, V]{
		keyToElement: make(map[K]*list.Element, capacity),
		linkedList:   list.New(),
		capacity:     capacity,
	}
}

func (c *lrucache[K, V]) Get(key K) (value V, ok bool) {
	if v, ok := c.keyToElement[key]; ok {
		c.linkedList.MoveToFront(v)
		return c.getNodeFromElement(v).value, true
	}

	var zeroValue V
	return zeroValue, false
}

func (c *lrucache[K, V]) Put(key K, value V) {
	if link, ok := c.keyToElement[key]; ok {
		n := c.getNodeFromElement(link)
		n.value = value
		c.linkedList.MoveToFront(link)
		return
	}

	if c.Size() == c.capacity {
		c.extractLatest()
	}

	c.keyToElement[key] = c.linkedList.PushFront(&node[K, V]{
		key:   key,
		value: value,
	})

}

func (c *lrucache[K, V]) Size() int {
	return len(c.keyToElement)
}

func (c *lrucache[K, V]) All() iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		cur := c.linkedList.Front()

		for range c.Size() {
			n := c.getNodeFromElement(cur)

			if !yield(n.key, n.value) {
				return
			}
			cur = cur.Next()
		}
	}
}

func (c *lrucache[K, V]) Clear() {
	c.linkedList = list.New()
	clear(c.keyToElement)
}

func (c *lrucache[K, V]) getNodeFromElement(element *list.Element) *node[K, V] {
	switch v := element.Value.(type) {
	case *node[K, V]:
		return v
	default:
		return nil
	}
}

func (c *lrucache[K, V]) extractLatest() {
	del := c.linkedList.Back()
	c.linkedList.Remove(del)
	delete(c.keyToElement, c.getNodeFromElement(del).key)
}
