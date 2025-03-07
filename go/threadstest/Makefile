.PHONY: run
run:
	go build && GOMAXPROCS=1 GODEBUG=schedtrace=1000,scheddetail=1 ./threadstest

.PHONY: run-gogs
run-gogs:
	go build && GOGC=off GOMAXPROCS=1 GODEBUG=schedtrace=1000,scheddetail=1 ./threadstest