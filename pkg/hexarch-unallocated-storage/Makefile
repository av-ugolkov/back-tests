.PHONY: build-docker
build-docker:
	docker build --tag kvs:multipart .

.PHONY: run-image
run-image:
	docker run --name kvs --detach --publish 8080:8080 kvs:multipart

.PHONY: gen
gen:
	protoc --proto_path=./proto \
	--go_out=./proto --go_opt=paths=source_relative \
	--go-grpc_out=./proto --go-grpc_opt=paths=source_relative \
	./proto/keyvalue.proto

.PHONY: curl-put
curl-put:
	curl -X PUT -d "key-value store" -v http://localhost:8080/v1/key-a

.PHONY: curl-get
curl-get:
	curl http://localhost:8080/v1/key-a
