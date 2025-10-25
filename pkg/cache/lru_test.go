package cache

import (
	"container/list"
	"reflect"
	"strings"
	"testing"
)

func Test_lrucache_Get(t *testing.T) {
	type args[K comparable] struct {
		key K
	}
	type testCase[K comparable, V any] struct {
		name      string
		c         lrucache[K, V]
		args      args[K]
		wantValue V
		wantOk    bool
	}

	tests := []testCase[int, string]{
		{name: testing.CoverMode(), c: struct {
			keyToElement map[int]*list.Element
			linkedList   *list.List
			capacity     int
		}{
			keyToElement: map[int]*list.Element{},
			linkedList:   list.New(),
			capacity:     5,
		},
			args:      struct{ key int }{key: 3},
			wantValue: "aaa",
			wantOk:    true,
		},
	}

	for _, tt := range tests {
		for i := 1; i <= tt.c.capacity; i++ {
			tt.c.Put(i, strings.Repeat("a", i))
		}

		t.Run(tt.name, func(t *testing.T) {
			gotValue, gotOk := tt.c.Get(tt.args.key)
			if !reflect.DeepEqual(gotValue, tt.wantValue) {
				t.Errorf("Get() gotValue = %v, want %v", gotValue, tt.wantValue)
			}
			if gotOk != tt.wantOk {
				t.Errorf("Get() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}

func Test_lrucache_Put(t *testing.T) {
	type args[K comparable, V any] struct {
		key   K
		value V
	}
	type testCase[K comparable, V any] struct {
		name      string
		c         lrucache[K, V]
		args      args[K, V]
		wantValue V
		wantOk    bool
	}
	tests := []testCase[int, string]{
		{name: testing.CoverMode(), c: struct {
			keyToElement map[int]*list.Element
			linkedList   *list.List
			capacity     int
		}{
			keyToElement: map[int]*list.Element{},
			linkedList:   list.New(),
			capacity:     5,
		},
			args: struct {
				key   int
				value string
			}{key: 3, value: "aaa"},
			wantValue: "aaa",
			wantOk:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.c.Put(tt.args.key, tt.args.value)

			value, ok := tt.c.Get(tt.args.key)
			if !reflect.DeepEqual(value, tt.wantValue) {
				t.Errorf("Get() gotValue = %v, want %v", value, tt.wantValue)
			}
			if ok != tt.wantOk {
				t.Errorf("Get() gotOk = %v, want %v", ok, tt.wantOk)
			}
		})
	}
}
