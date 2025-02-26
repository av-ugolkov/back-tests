build:
	@go build -buildmode=plugin -o duck/duck.so duck/main.go
	@go build -buildmode=plugin -o frog/frog.so frog/main.go

frog: build
	@go run main/main.go frog

duck: build
	@go run main/main.go duck
