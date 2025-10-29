.PHONY: zookeeper
zookeeper:
	docker compose -p backend-examples -f common.yml -f zookeeper.yml up

.PHONY: kafka_cluster
kafka_cluster:
	docker compose -p backend-examples -f common.yml -f kafka_cluster.yml up

.PHONY: kafka_init
kafka_init:
	docker compose -p backend-examples -f common.yml -f kafka_init.yml up

.PHONY: show-kafka-test-request
show-kafka-test-request:
	kafkacat -C -b localhost:19092 -t kafka-test-request