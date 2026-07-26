package main

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"

	"github.com/uptrace/bunrouter"
)

func bunrouterFramework() framework {
	return framework{
		name: "bunrouter",
		serve: func(ln net.Listener) func() {
			r := bunrouter.New()

			r.GET("/ping", func(w http.ResponseWriter, req bunrouter.Request) error {
				w.WriteHeader(http.StatusOK)
				return nil
			})

			r.GET("/users/:id", func(w http.ResponseWriter, req bunrouter.Request) error {
				u := sampleUser
				u.ID = req.Param("id")
				return bunrouter.JSON(w, u)
			})

			r.POST("/users", func(w http.ResponseWriter, req bunrouter.Request) error {
				var u User
				if err := json.NewDecoder(req.Body).Decode(&u); err != nil {
					return fmt.Errorf("decode user: %w", err)
				}
				return bunrouter.JSON(w, u)
			})

			return serveHTTP(ln, r)
		},
	}
}
