package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/Shopify/sarama"
)

type ProducerConfig struct {
	kafkaHost  string
	kafkaPort  string
	kafkaTopic string
}

var pubConfig = ProducerConfig{
	kafkaHost:  getEnv("KAFKA_HOST", "broker.kafka"),
	kafkaPort:  getEnv("KAFKA_PORT", "9092"),
	kafkaTopic: getEnv("KAFKA_TOPIC", "banana"),
}

func main() {
	// create producer
	producer, err := initProducer()
	if err != nil {
		fmt.Println("Error producer: ", err.Error())
		os.Exit(1)
	}

	// read command line input
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter msg: ")
		msg, _ := reader.ReadString('\n')

		// publish without goroutine
		publish(msg, producer)

		// publish with go routine
		// go publish(msg, producer)
	}
}

func initProducer() (sarama.SyncProducer, error) {
	broker := pubConfig.kafkaHost + ":" + pubConfig.kafkaPort

	// setup sarama log to stdout
	sarama.Logger = log.New(os.Stdout, "", log.Ltime)

	// producer config
	config := sarama.NewConfig()
	config.Producer.Retry.Max = 5
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true

	// async producer
	//prd, err := sarama.NewAsyncProducer([]string{kafkaConn}, config)

	// sync producer
	prd, err := sarama.NewSyncProducer([]string{broker}, config)

	return prd, err
}

func publish(message string, producer sarama.SyncProducer) {
	// publish sync
	msg := &sarama.ProducerMessage{
		Topic: pubConfig.kafkaTopic,
		Value: sarama.StringEncoder(message),
	}
	p, o, err := producer.SendMessage(msg)
	if err != nil {
		fmt.Println("Error publish: ", err.Error())
	}

	// publish async
	//producer.Input() <- &sarama.ProducerMessage{

	fmt.Println("Partition: ", p)
	fmt.Println("Offset: ", o)
}

func getEnv(key, notset string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return notset
}
