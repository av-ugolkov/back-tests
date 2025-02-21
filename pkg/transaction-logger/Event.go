package transactionlogger

type EventType byte

const (
	EventDelete EventType = 1
	EventPut    EventType = 2
)

type Event struct {
	Sequence  uint64
	EventType EventType
	Key       string
	Value     string
}
