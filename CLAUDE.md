# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## What this repository is

A personal learning/reference repository of standalone Go examples, not a deployable application. Module `github.com/av-ugolkov/backend-examples` (Go 1.26). Roughly 90+ directories each contain their own `main.go` with `package main` — almost every leaf directory is an independent runnable program. There is no root entrypoint, no root Makefile, and no CI.

`for_deleting/`, `leet-code/`, `temp/` are gitignored scratch areas — ignore them unless explicitly asked.

## Commands

Examples are run per-directory, not from the repo root:

```bash
go run ./go/goroutine/gosched          # run one example
go test ./go/strings/split/            # test one package
go test -run TestName ./pkg/cache/     # single test
go test -bench=. ./go/http/echo/       # benchmarks (echo vs gin comparison lives under go/http/)
go vet ./...                           # whole-repo check — this is clean and is the correct repo-wide gate
go test ./...                          # also clean
```

**Do not use `go build ./...` as a health check.** Several directories declare `package main` without a `main()` (`solid/*`, `go/plugin/duck`, `go/plugin/frog`) and fail at link time by design. Use `go vet ./...` instead.

Many subdirectories carry a local `Makefile` for their specific workflow — check for one before inventing a command. Recurring targets:

- `make gen` / `make gen-proto` — regenerate protobuf/gRPC stubs (`go/gRPC`, `go/protobuf`, both `*unallocated-storage` dirs). Requires `protoc` + Go plugins.
- `make run` / `make docker-run` / `make database` — `docker compose up` for examples needing infrastructure (`pkg/goobs`, `pkg/kafka`, `go/metric/otel`, `db/postgresql/*`).
- `make build-docker` / `make run-image` — container build for the key-value store examples.
- k8s examples (`k8s/*/Makefile`) target minikube: `docker.build`, `minikube.build`, `k8s.apply`.

Special build modes: `go/cgo/*` needs `CGO_ENABLED=1`; `go/plugin` requires `-buildmode=plugin` (`make build` first); `go/build/tags` toggles behavior via `-tags debug`; `go/data-race` is meant to be run with `-race`.

## Layout

- **`go/`** — language and stdlib feature demos, one concept per directory (goroutines, defer, generics, syscall, pprof, weak pointers, cgo, plugins, otel trace/metrics, logging with slog/zap, http router comparisons).
- **`patterns/`** — concurrency patterns (fan-in/out, pipeline, semaphore, circuit-breaker, worker-pool, single-flight, …). Each is `main.go` (driver) + a file named after the pattern holding the generic implementation.
- **`pkg/`** — the only multi-file, architecture-bearing projects; see below.
- **`solid/`** — SOLID principle examples, ported from `github.com/MaksimDzhangirov/practicalSolid`. Note `solid/dip/infrastructire` is a typo in the actual path.
- **`db/postgresql/`**, **`k8s/`** — SQL scripts and manifests with docker-compose/minikube Makefiles; little to no Go.

## The `pkg/` projects

**`unallocated-storage` → `hexarch-unallocated-storage`** are the same key-value store at two stages of refactoring; changes to one are often meant for both. Both expose PUT/GET/DELETE on `/v1/{key}` at `:8080` via `gorilla/mux`, plus a gRPC `server`/`client` pair.

- `unallocated-storage` is the flat version: package-level state in `core`, handlers wired directly in `handler.go`, transaction logging pulled from the shared `pkg/transaction-logger`.
- `hexarch-unallocated-storage` is the hexagonal version. `core` owns the domain (`KeyValueStore` + the `TransactionLogger` port) and depends on nothing else. Adapters are selected at runtime by env var through factories: `FRONTEND_TYPE` (`rest`|`zero`) in `frontend/factory.go`, `TLOG_TYPE` (`file`|`postgres`, with `TLOG_FILENAME`) in `transact/factory.go`. Adding an adapter means implementing the port and adding a case to the factory — `core` must not import `frontend` or `transact`.

**`pkg/transaction-logger`** — write-ahead event log (file and postgres backends) used by the flat KVS; the hexagonal one has its own copy under `transact/`.

**`pkg/goobs`** — **a separate nested Go module** (`module goobs`, its own `go.mod`). Root-level `go` commands do not cover it; `cd pkg/goobs` first. Fiber service instrumented with OpenTelemetry tracing (Jaeger) and Prometheus metrics; `make run` brings up the whole app+Grafana+Prometheus+Jaeger stack.

**`pkg/kafka`** — producer/consumer against a 3-broker cluster on `localhost:19092,29092,39092`, built on `confluent-kafka-go` (cgo/librdkafka). Compose files are split and layered: `make zookeeper`, `make kafka_cluster`, `make kafka_init`, each running `docker compose -f common.yml -f <file>.yml`. `pkg/kafka/volumes/` is gitignored state.

**`pkg/cache`** — standalone LRU implementation, one of the few genuinely library-style packages here.

## Conventions in this codebase

- New examples follow the existing shape: a new leaf directory, `package main`, a `main.go` that demonstrates the concept and prints results — not a library package unless it belongs in `pkg/`.
- Generic constraints are used freely in `patterns/` (see `ChanType` in `patterns/fan-in`); prefer generics over `interface{}`.
- Generated protobuf files (`*.pb.go`) are committed; regenerate via the local `make gen` rather than editing them.
- Some existing comments are in Russian. Match the user's global instruction — minimal comments, doc comments only — for new code.