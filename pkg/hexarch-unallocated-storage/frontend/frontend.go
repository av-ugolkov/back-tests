package frontend

import "github.com/av-ugolkov/backend-examples/pkg/hexarch-unallocated-storage/core"

type FrontEnd interface {
	Start(kv *core.KeyValueStore) error
}

type zeroFrontEnd struct{}

func (f zeroFrontEnd) Start(kv *core.KeyValueStore) error {
	return nil
}
