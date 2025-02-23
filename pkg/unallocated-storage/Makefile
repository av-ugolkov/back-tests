.PHONY: build-docker
build-docker:
	docker build --tag kvs:multipart .

.PHONY: run-image
run-image:
	docker run --name kvs --detach --publish 8080:8080 kvs:multipart