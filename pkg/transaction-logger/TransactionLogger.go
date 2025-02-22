package transactionlogger

type TransactionLogger interface {
	Close() error
	WriteDelete(key string)
	WritePut(key, value string)
	Err() <-chan error
	ReadEvents() (<-chan Event, <-chan error)
	Run()
}
