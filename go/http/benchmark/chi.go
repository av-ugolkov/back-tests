package main

import (
	"encoding/json"
	"net"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func chiFramework() framework {
	return framework{
		name: "chi",
		serve: func(ln net.Listener) func() {
			r := chi.NewRouter()

			r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
			})

			r.Get("/users/{id}", func(w http.ResponseWriter, r *http.Request) {
				u := sampleUser
				u.ID = chi.URLParam(r, "id")
				writeJSON(w, u)
			})

			r.Post("/users", func(w http.ResponseWriter, r *http.Request) {
				var u User
				if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
					http.Error(w, err.Error(), http.StatusBadRequest)
					return
				}
				writeJSON(w, u)
			})

			return serveHTTP(ln, r)
		},
	}
}
