.PHONY: build-docker
build-docker:
	docker build --tag kvs:multipart .

.PHONY: run-image
run-image:
	docker run --name kvs --detach --publish 8080:8080 kvs:multipart

.PHONY: curl-put
curl-put:
	curl -X PUT -d "key-value store" -v http://localhost:8080/v1/key-a

.PHONY: curl-get
curl-get:
	curl http://localhost:8080/v1/key-a