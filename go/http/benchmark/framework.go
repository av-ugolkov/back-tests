package main

import (
	"errors"
	"log"
	"net"
	"net/http"
)

// framework is one router under test. serve starts it on ln and returns a
// function that shuts it down; ln is closed by the shutdown.
type framework struct {
	name  string
	serve func(net.Listener) (stop func())
}

func frameworks() []framework {
	return []framework{
		stdlibFramework(),
		ginFramework(),
		echoFramework(),
		fiberFramework(),
		bunrouterFramework(),
		chiFramework(),
	}
}

// serveHTTP runs h behind an identical net/http server for every framework
// that is not fasthttp-based, so the transport layer is not a variable.
func serveHTTP(ln net.Listener, h http.Handler) func() {
	srv := &http.Server{Handler: h}

	go func() {
		if err := srv.Serve(ln); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Printf("serve: %v", err)
		}
	}()

	return func() {
		if err := srv.Close(); err != nil {
			log.Printf("close server: %v", err)
		}
	}
}
