.PHONY: docker.build
docker.build:
	docker build . -f deploy/Dockerfile -t app:v1

.PHONY: minikube.start
minikube.start:
	minikube start -p goapp --driver=docker --extra-config=apiserver.service-node-port-range=1000-50000

.PHONY: minikube.load
minikube.load:
	minikube -p goapp image load app:v2

.PHONY: minikube.build
minikube.build:
	minikube -p goapp image build . -t app:v2 -f deploy/Dockerfile

.PHONY: minikube.tunnel
minikube.tunnel:
	minikube -p goapp -n develop service goapp-svc --url

.PHONY: k8s.apply
k8s.apply:
	kubectl apply -f deploy/dev.yaml

.PHONY: k8s.forward
k8s.forward:
	kubectl -n develop port-forward service/goapp-svc 5858:80