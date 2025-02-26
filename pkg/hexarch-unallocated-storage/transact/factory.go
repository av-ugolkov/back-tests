package transact

import (
	"fmt"
	"os"

	"github.com/av-ugolkov/backend-examples/pkg/hexarch-unallocated-storage/core"
)

func NewTransactionLogger(logger string) (core.TransactionLogger, error) {
	switch logger {
	case "file":
		return NewFileLogger(os.Getenv("TLOG_FILENAME"))
	case "postgres":
		return NewPostgresLogger(PostgresDBParams{
			host:     "localhost",
			dbName:   "kvs",
			user:     "test",
			password: "hunter2",
		})
	case "":
		return nil, fmt.Errorf("transaction logger type not defined")
	default:
		return nil, fmt.Errorf("no such transaction logger %s", logger)
	}
}
