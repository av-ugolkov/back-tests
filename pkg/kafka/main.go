package main

import (
	"context"
	"fmt"
	"math/rand/v2"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

const topicName = "kafka-test-request"

func main() {
	cfgProducer := kafka.ConfigMap{
		"bootstrap.servers":      "localhost:19092,localhost:29092,localhost:39092",
		"client.id":              "backend-examples",
		"acks":                   "all",
		"linger.ms":              0,
		"batch.num.messages":     1,
		"queue.buffering.max.ms": 0,
	}
	producer := NewKafkaProducer(&cfgProducer)

	cfgConsumer := kafka.ConfigMap{
		"bootstrap.servers":  "localhost:19092,localhost:29092,localhost:39092",
		"group.id":           "backend-examples",
		"auto.offset.reset":  "earliest",
		"enable.auto.commit": false,
	}
	consumer := NewKafkaConsumer(&cfgConsumer)

	r := NewReceiver()
	consumer.SubscribeTopics(topicName, r)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		err := consumer.Listen(ctx, []string{topicName})
		if err != nil {
			fmt.Printf("error listen: %v\n", err)
		}
	}()

	var close bool
	go func() {
		for !close {
			ind := time.Now().Unix()
			err := producer.Send(topicName, []byte(fmt.Sprintf("%d", ind)), []byte(fmt.Sprintf("%v", time.Now().Format(time.DateTime))))
			if err != nil {
				fmt.Printf("error producer send: %v\n", err)
			}
			wait := rand.IntN(2500) + 500
			time.Sleep(time.Duration(wait) * time.Millisecond)
		}
		producer.Close()
	}()

	sig := <-sigchan
	fmt.Printf("caught signal [%v]: terminating\n", sig)
	close = true
	time.Sleep(3 * time.Second)
	consumer.Close()
	time.Sleep(3 * time.Second)
}
