package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"

	transactionlogger "github.com/av-ugolkov/backend-examples/pkg/transaction-logger"
	"github.com/av-ugolkov/backend-examples/pkg/unallocated-storage/core"

	"github.com/gorilla/mux"
)

func Init() {
	err := initializeTransactionLog()
	if err != nil {
		log.Fatal(err)
	}

	r := mux.NewRouter()

	r.HandleFunc("/v1/{key}", keyValuePutHandler).Methods("PUT")
	r.HandleFunc("/v1/{key}", keyValueGetHandler).Methods("GET")
	r.HandleFunc("/v1/{key}", keyValueDeleteHandler).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", r))
}

var logger transactionlogger.TransactionLogger

func initializeTransactionLog() error {
	var err error

	logger, err = transactionlogger.New("transaction.log")
	if err != nil {
		return fmt.Errorf("failed to create event logger: %w", err)
	}

	chEvents, chErrors := logger.ReadEvents()
	e, ok := transactionlogger.Event{}, true

	for ok && err == nil {
		select {
		case err, ok = <-chErrors:
		case e, ok = <-chEvents:
			switch e.EventType {
			case transactionlogger.EventDelete:
				err = core.Delete(e.Key)
			case transactionlogger.EventPut:
				err = core.Put(e.Key, e.Value)
			}
		}
	}

	logger.Run()

	return err
}

func keyValuePutHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]

	value, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = core.Put(key, string(value))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	logger.WritePut(key, string(value))

	w.WriteHeader(http.StatusCreated)
}

func keyValueGetHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]

	value, err := core.Get(key)
	if errors.Is(err, core.ErrorNoSuchKey) {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte(value))
}

func keyValueDeleteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]

	err := core.Delete(key)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	logger.WriteDelete(key)

	w.WriteHeader(http.StatusOK)
}
