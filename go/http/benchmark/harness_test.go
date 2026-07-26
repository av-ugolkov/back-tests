package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net"
	"net/http"

	"github.com/valyala/fasthttp/fasthttputil"
)

type benchCase struct {
	name   string
	method string
	path   string
	body   []byte
}

func benchCases() []benchCase {
	return []benchCase{
		{name: "status", method: http.MethodGet, path: "/ping"},
		{name: "json-out", method: http.MethodGet, path: "/users/42"},
		{name: "json-io", method: http.MethodPost, path: "/users", body: sampleUserJSON},
	}
}

// harness drives one framework over an in-memory listener, so no TCP stack or
// syscall cost enters the measurement and every framework is reached through
// the same http.Client.
type harness struct {
	client *http.Client
	stop   func()
}

func start(f framework) *harness {
	ln := fasthttputil.NewInmemoryListener()
	stopServer := f.serve(ln)

	client := &http.Client{
		Transport: &http.Transport{
			DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
				return ln.Dial()
			},
			MaxIdleConns:        1024,
			MaxIdleConnsPerHost: 1024,
			DisableCompression:  true,
		},
	}

	return &harness{
		client: client,
		stop: func() {
			client.CloseIdleConnections()
			stopServer()
		},
	}
}

func (h *harness) newRequest(c benchCase) (*http.Request, error) {
	var body io.Reader
	if c.body != nil {
		body = bytes.NewReader(c.body)
	}

	req, err := http.NewRequest(c.method, "http://inmemory"+c.path, body)
	if err != nil {
		return nil, fmt.Errorf("build request: %w", err)
	}
	if c.body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	return req, nil
}

// do runs one request and discards the body. Draining and closing keeps the
// connection in the idle pool; skipping it would make every iteration pay for
// a fresh dial.
func (h *harness) do(c benchCase) (int, error) {
	req, err := h.newRequest(c)
	if err != nil {
		return 0, err
	}

	resp, err := h.client.Do(req)
	if err != nil {
		return 0, fmt.Errorf("do request: %w", err)
	}

	if _, err := io.Copy(io.Discard, resp.Body); err != nil {
		resp.Body.Close()
		return resp.StatusCode, fmt.Errorf("drain body: %w", err)
	}
	if err := resp.Body.Close(); err != nil {
		return resp.StatusCode, fmt.Errorf("close body: %w", err)
	}

	return resp.StatusCode, nil
}

func (h *harness) doBody(c benchCase) (int, []byte, error) {
	req, err := h.newRequest(c)
	if err != nil {
		return 0, nil, err
	}

	resp, err := h.client.Do(req)
	if err != nil {
		return 0, nil, fmt.Errorf("do request: %w", err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		resp.Body.Close()
		return resp.StatusCode, nil, fmt.Errorf("read body: %w", err)
	}
	if err := resp.Body.Close(); err != nil {
		return resp.StatusCode, body, fmt.Errorf("close body: %w", err)
	}

	return resp.StatusCode, body, nil
}
