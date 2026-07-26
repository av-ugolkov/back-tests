package main

import (
	"encoding/json"
	"log"
	"net"
	"net/http"
)

func stdlibFramework() framework {
	return framework{
		name: "stdlib",
		serve: func(ln net.Listener) func() {
			mux := http.NewServeMux()

			mux.HandleFunc("GET /ping", func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
			})

			mux.HandleFunc("GET /users/{id}", func(w http.ResponseWriter, r *http.Request) {
				u := sampleUser
				u.ID = r.PathValue("id")
				writeJSON(w, u)
			})

			mux.HandleFunc("POST /users", func(w http.ResponseWriter, r *http.Request) {
				var u User
				if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
					http.Error(w, err.Error(), http.StatusBadRequest)
					return
				}
				writeJSON(w, u)
			})

			return serveHTTP(ln, mux)
		},
	}
}

func writeJSON(w http.ResponseWriter, v any) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(v); err != nil {
		log.Printf("stdlib encode: %v", err)
	}
}
