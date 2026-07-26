package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
)

// TestRoutes guards the benchmark: a framework whose routes are wired wrong
// answers 404 faster than a correct one does and would top the table.
func TestRoutes(t *testing.T) {
	for _, f := range frameworks() {
		t.Run(f.name, func(t *testing.T) {
			h := start(f)
			t.Cleanup(h.stop)

			for _, c := range benchCases() {
				t.Run(c.name, func(t *testing.T) {
					code, body, err := h.doBody(c)
					if err != nil {
						t.Fatalf("request: %v", err)
					}
					if code != http.StatusOK {
						t.Fatalf("expected 200, got %d (body %q)", code, body)
					}

					if c.name == "status" {
						if len(bytes.TrimSpace(body)) != 0 {
							t.Fatalf("expected empty body, got %q", body)
						}
						return
					}

					var got User
					if err := json.Unmarshal(body, &got); err != nil {
						t.Fatalf("unmarshal %q: %v", body, err)
					}
					if got != sampleUser {
						t.Fatalf("expected %+v, got %+v", sampleUser, got)
					}
				})
			}
		})
	}
}
