package main

import (
	"log"
	"os"

	"github.com/av-ugolkov/backend-examples/pkg/hexarch-unallocated-storage/core"
	"github.com/av-ugolkov/backend-examples/pkg/hexarch-unallocated-storage/frontend"
	"github.com/av-ugolkov/backend-examples/pkg/hexarch-unallocated-storage/transact"
)

func main() {
	tl, err := transact.NewTransactionLogger(os.Getenv("TLOG_TYPE"))
	if err != nil {
		log.Fatal(err)
	}

	store := core.NewKeyValueStore(tl)
	store.Restore()

	fe, err := frontend.NewFrontEnd(os.Getenv("FRONTEND_TYPE"))
	if err != nil {
		log.Fatal(err)
	}

	log.Fatal(fe.Start(store))
}
