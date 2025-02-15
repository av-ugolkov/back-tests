.PHONY: run
run:
	CGO_ENABLED=1 go run main.go

.PHONY: build_sum
build_sum:
	gcc -shared -o sum/libsum.so -fPIC sum/src/sum.c