package main

import (
	"bytes"
	"encoding/json"
	"sync"
)

type User struct {
	Name string
	Age  int
}

var buf *bytes.Buffer
var pool = sync.Pool{
	New: func() any {
		return json.NewEncoder(buf)
	},
}

func main() {
	user := User{"Alice", 30}
	withoutPool(user)
	buf = bytes.NewBuffer(make([]byte, 0, 100))
	withPool(user)
}

func withoutPool(u User) []byte {
	data, _ := json.Marshal(u)
	return data
}

func withPool(u User) []byte {
	encoder := pool.Get().(*json.Encoder)
	encoder.Encode(&u)
	pool.Put(encoder)
	return buf.Bytes()
}
