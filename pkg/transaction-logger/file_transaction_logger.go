package transactionlogger

import (
	"bufio"
	"fmt"
	"os"
)

/*TODO
отсутствуют тесты;
нет метода Close для корректного закрытия файла;
служба может закрыться, когда события все еще находятся в буфере
записи: в этом случае события могут быть потеряны;
ключи и значения никак не кодируются в журнале транзакций: если событие будет содержать несколько строк или пробелов, то его не удастся
правильно проанализировать;
размеры ключей и значений не ограничены: служба позволяет добавлять огромные ключи или значения, что может привести к переполнению диска;
журнал транзакций хранится в простом текстовом виде: он будет занимать больше места на диске, чем мог бы;
журнал хранит записи об удаленных значениях: он будет неограниченно расти.
*/

type FileTransactionLogger struct {
	events       chan<- Event
	errors       <-chan error
	lastSequence uint64
	file         *os.File
}

func NewFileLogger(filename string) (TransactionLogger, error) {
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0755)
	if err != nil {
		return nil, fmt.Errorf("cannot open transaction log file: %w", err)
	}

	return &FileTransactionLogger{file: file}, nil
}

func (l *FileTransactionLogger) Close() error {
	err := l.file.Close()
	if err != nil {
		return fmt.Errorf("cannot close transaction log file: %w", err)
	}

	return nil
}

func (l *FileTransactionLogger) Run() {
	events := make(chan Event, 16)
	l.events = events

	errors := make(chan error, 1)
	l.errors = errors

	go func() {
		for e := range events {
			l.lastSequence++

			_, err := fmt.Fprintf(l.file, "%d\t%d\t%s\t%s\n", l.lastSequence, e.EventType, e.Key, e.Value)
			if err != nil {
				errors <- err
				return
			}
		}
	}()
}

func (l *FileTransactionLogger) ReadEvents() (<-chan Event, <-chan error) {
	scanner := bufio.NewScanner(l.file)
	outEvent := make(chan Event)
	outError := make(chan error, 1)

	go func() {
		var e Event

		defer close(outEvent)
		defer close(outError)

		for scanner.Scan() {
			line := scanner.Text()

			if _, err := fmt.Sscanf(line, "%d\t%d\t%s\t%s",
				&e.Sequence,
				e.EventType,
				e.Key,
				e.Value); err != nil {
				outError <- fmt.Errorf("transaction numbers out of sequence")
				return
			}

			l.lastSequence = e.Sequence
			outEvent <- e
		}

		if err := scanner.Err(); err != nil {
			outError <- fmt.Errorf("transaction log read failure: %w", err)
			return
		}
	}()

	return outEvent, outError
}

func (l *FileTransactionLogger) WriteDelete(key string) {
	l.events <- Event{EventType: EventDelete, Key: key}
}

func (l *FileTransactionLogger) WritePut(key, value string) {
	l.events <- Event{EventType: EventPut, Key: key, Value: value}
}

func (l *FileTransactionLogger) Err() <-chan error {
	return l.errors
}
