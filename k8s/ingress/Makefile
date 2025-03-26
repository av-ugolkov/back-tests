.PHONY: docker.build
docker.build:
	docker build . -f deploy/Dockerfile -t app:v1

.PHONY: minikube.start
minikube.start:
	minikube start --driver=docker

.PHONY: minikube.load
minikube.load:
	minikube image load app:v1

.PHONY: minikube.build
minikube.build:
	minikube image build . -t app:v1 -f deploy/Dockerfile

.PHONY: minikube.srv
minikube.srv:
	minikube -n develop service goapp-svc --url

.PHONY: minikube.tunnel
minikube.tunnel:
	minikube tunnel

.PHONY: k8s.apply
k8s.apply:
	kubectl apply -f deploy/dev.yaml