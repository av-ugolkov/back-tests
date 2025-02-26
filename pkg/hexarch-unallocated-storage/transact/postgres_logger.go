package transact

import (
	"database/sql"
	"fmt"

	"github.com/av-ugolkov/backend-examples/pkg/hexarch-unallocated-storage/core"
)

/*TODO
предполагается, что база данных и таблица существуют, и код потерпит
неудачу, если это не так;
строка подключения жестко запрограммирована, даже пароль;
по-прежнему отсутствует метод Close, который закрывал бы соединение;
служба может закрыться, когда события все еще находятся в буфере
записи: в этом случае события могут быть потеряны;
журнал хранит записи об удаленных значениях: он будет неограниченно расти.
*/

type PostgresDBParams struct {
	dbName   string
	host     string
	user     string
	password string
}

type PostgresTransactionLogger struct {
	events chan<- core.Event
	errors <-chan error
	db     *sql.DB
}

func NewPostgresLogger(cfg PostgresDBParams) (core.TransactionLogger, error) {
	connStr := fmt.Sprintf("host=%s dbname=%s user=%s password=%s",
		cfg.host, cfg.dbName, cfg.user, cfg.password)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to open db: %w", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed to open db connection: %w", err)
	}

	logger := &PostgresTransactionLogger{db: db}

	exists, err := logger.verifyTableExists()
	if err != nil {
		return nil, fmt.Errorf("failed to verify table exists: %w", err)
	}
	if !exists {
		if err = logger.createTable(); err != nil {
			return nil, fmt.Errorf("failed to create table: %w", err)
		}
	}

	return logger, nil
}

func (l *PostgresTransactionLogger) verifyTableExists() (bool, error) {
	var count int
	err := l.db.QueryRow("SELECT count(*) FROM INFORMATION_SCHEMA.TABLES WHERE TABLE_NAME='transactions'").Scan(count)
	if err != nil {
		return false, fmt.Errorf("failed to exec command in postgres: %w", err)
	}

	return count == 1, nil
}

func (l *PostgresTransactionLogger) createTable() error {
	query := `
		CREATE TABLE IF NOT EXISTS "transactions"(
			sequence      BIGSERIAL PRIMARY KEY,
			event_type    SMALLINT,
			key 		  TEXT,
			value         TEXT
		);`

	_, err := l.db.Exec(query)
	if err != nil {
		return fmt.Errorf("failed to exec command in postgres: %w", err)
	}

	return nil
}

func (l *PostgresTransactionLogger) Close() error {
	return l.db.Close()
}

func (l *PostgresTransactionLogger) WritePut(key, value string) {
	l.events <- core.Event{EventType: core.EventPut, Key: key, Value: value}
}

func (l *PostgresTransactionLogger) WriteDelete(key string) {
	l.events <- core.Event{EventType: core.EventDelete, Key: key}
}

func (l *PostgresTransactionLogger) Err() <-chan error {
	return l.errors
}

func (l *PostgresTransactionLogger) ReadEvents() (<-chan core.Event, <-chan error) {
	outEvent := make(chan core.Event)
	outError := make(chan error, 1)

	go func() {
		defer close(outEvent)
		defer close(outError)

		query := `
			SELECT sequence, event_type, key, value
			FROM transactions
			ORDER BY sequence;`

		rows, err := l.db.Query(query)
		if err != nil {
			outError <- fmt.Errorf("sql query error: %w", err)
			return
		}
		defer rows.Close()

		e := core.Event{}

		for rows.Next() {
			err = rows.Scan(&e.Sequence, &e.EventType, &e.Key, &e.Value)
			if err != nil {
				outError <- fmt.Errorf("error reading row: %w", err)
				return
			}

			outEvent <- e
		}

		err = rows.Err()
		if err != nil {
			outError <- fmt.Errorf("transaction log read failure: %w", err)
		}
	}()
	return outEvent, outError
}

func (l *PostgresTransactionLogger) Run() {
	events := make(chan core.Event, 16)
	l.events = events

	errors := make(chan error, 1)
	l.errors = errors

	go func() {
		query := `INSERT INTO transactions(event_type, key, value)
			VALUES ($1, $2, $3)`

		for e := range events {
			_, err := l.db.Exec(query, e.EventType, e.Key, e.Value)
			if err != nil {
				errors <- err
			}
		}
	}()
}
