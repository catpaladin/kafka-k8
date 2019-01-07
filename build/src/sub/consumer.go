package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Shopify/sarama"
	"github.com/wvanbergen/kafka/consumergroup"
)

type ConsumerConfig struct {
	zookeeperHost string
	zookeeperPort string
	consumerGroup string
	kafkaTopic    string
}

var subConfig = ConsumerConfig{
	zookeeperHost: getEnv("ZOOKEEPER_HOST", "zookeeper.kafka"),
	zookeeperPort: getEnv("ZOOKEEPER_PORT", "2181"),
	consumerGroup: getEnv("CONSUMER_GROUP", "zbanana"),
	kafkaTopic:    getEnv("KAFKA_TOPIC", "banana"),
}

func main() {
	// setup sarama log to stdout
	sarama.Logger = log.New(os.Stdout, "", log.Ltime)

	// init consumer
	cg, err := initConsumer()
	if err != nil {
		fmt.Println("Error consumer goup: ", err.Error())
		os.Exit(1)
	}
	defer cg.Close()

	// run consumer
	consume(cg)
}

func initConsumer() (*consumergroup.ConsumerGroup, error) {
	zoo := subConfig.zookeeperHost + ":" + subConfig.zookeeperPort

	// consumer config
	config := consumergroup.NewConfig()
	config.Offsets.Initial = sarama.OffsetOldest
	config.Offsets.ProcessingTimeout = 10 * time.Second

	// join to consumer group
	cg, err := consumergroup.JoinConsumerGroup(subConfig.consumerGroup, []string{subConfig.kafkaTopic}, []string{zoo}, config)
	if err != nil {
		return nil, err
	}

	return cg, err
}

func consume(cg *consumergroup.ConsumerGroup) {
	for {
		select {
		case msg := <-cg.Messages():
			// messages coming through chanel
			// only take messages from subscribed topic
			if msg.Topic != subConfig.kafkaTopic {
				continue
			}

			fmt.Println("Topic: ", msg.Topic)
			fmt.Println("Value: ", string(msg.Value))

			// commit to zookeeper that message is read
			// this prevent read message multiple times after restart
			err := cg.CommitUpto(msg)
			if err != nil {
				fmt.Println("Error commit zookeeper: ", err.Error())
			}
		}
	}
}

func getEnv(key, notset string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return notset
}
