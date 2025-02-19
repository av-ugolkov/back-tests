package main

import (
	"fmt"
)

func main() {
	shardedMap := New[string](5)

	shardedMap.Set("key1", "value1")
	shardedMap.Set("key2", "value2")
	shardedMap.Set("key3", "value3")

	fmt.Println(shardedMap.Get("key1"))
	fmt.Println(shardedMap.Get("key2"))
	fmt.Println(shardedMap.Get("key3"))

	shardedMap.Delete("key2")
	fmt.Println(shardedMap.Contains("key2"))

	keys := shardedMap.Keys()
	fmt.Println(keys)
}
