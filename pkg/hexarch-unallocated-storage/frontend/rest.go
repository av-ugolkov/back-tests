package frontend

import (
	"net/http"

	"github.com/av-ugolkov/backend-examples/pkg/hexarch-unallocated-storage/core"
	"github.com/gorilla/mux"
)

type restFrontEnd struct {
	store *core.KeyValueStore
}

func (f *restFrontEnd) Start(store *core.KeyValueStore) error {
	f.store = store

	r := mux.NewRouter()
	r.HandleFunc("/v1/{key}", f.keyValueGetHandler).Methods("GET")
	r.HandleFunc("/v1/{key}", f.keyValuePutHandler).Methods("PUT")
	r.HandleFunc("/v1/{key}", f.keyValueDeleteHandler).Methods("DELETE")

	return http.ListenAndServe(":8080", r)
}

func (f *restFrontEnd) keyValueGetHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]

	value, err := f.store.Get(key)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(value))
}

func (f *restFrontEnd) keyValuePutHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]
	value := vars["value"]

	err := f.store.Put(key, value)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (f *restFrontEnd) keyValueDeleteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]

	err := f.store.Delete(key)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
