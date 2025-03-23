.PHONY: docker.build
docker.build:
	docker build . -f deploy/Dockerfile -t app:v1

.PHONY: minikube.build
minikube.build:
	minikube image build . -t app:v1 -f deploy/Dockerfile

.PHONY: k8s.apply
k8s.apply:
	kubectl apply -f deploy/dev.yaml