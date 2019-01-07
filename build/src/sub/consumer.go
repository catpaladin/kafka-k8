package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/Shopify/sarama"
)

type ConsumerConfig struct {
	kafkaHost  string
	kafkaPort  string
	kafkaTopic string
}

var subConfig = ConsumerConfig{
	kafkaHost:  getEnv("KAFKA_HOST", "broker.kafka"),
	kafkaPort:  getEnv("KAFKA_PORT", "9092"),
	kafkaTopic: getEnv("KAFKA_TOPIC", "banana"),
}

func main() {
	broker := subConfig.kafkaHost + ":" + subConfig.kafkaPort

	// Simple sarama consumer
	// https://godoc.org/github.com/Shopify/sarama#example-Consumer
	consumer, err := sarama.NewConsumer([]string{broker}, nil)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := consumer.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	partitionConsumer, err := consumer.ConsumePartition(subConfig.kafkaTopic, 0, sarama.OffsetOldest)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := partitionConsumer.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	// Trap SIGINT to trigger a shutdown.
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	consumed := 0
ConsumerLoop:
	for {
		select {
		case msg := <-partitionConsumer.Messages():
			log.Printf("Consumed message offset %d\n", msg.Offset)
			log.Printf("Message: %s\n", msg.Value)
			consumed++
		case <-signals:
			break ConsumerLoop
		}
	}

	log.Printf("Consumed: %d\n", consumed)
}

func getEnv(key, notset string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return notset
}
