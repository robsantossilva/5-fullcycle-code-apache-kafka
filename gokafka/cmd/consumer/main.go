package main

//kafka-consumer-groups --bootstrap-server=localhost:9092 --describe --group=goapp-group

// GROUP           TOPIC           PARTITION  CURRENT-OFFSET  LOG-END-OFFSET  LAG             CONSUMER-ID                                         HOST            CLIENT-ID
// goapp-group     teste           0          7               7               0               goapp-consumer-f82c55ed-27bf-4933-9673-da86bfe1da75 /172.27.0.3     goapp-consumer
// goapp-group     teste           1          9               9               0               goapp-consumer-f82c55ed-27bf-4933-9673-da86bfe1da75 /172.27.0.3     goapp-consumer
// goapp-group     teste           2          25              25              0               goapp-consumer-f82c55ed-27bf-4933-9673-da86bfe1da75 /172.27.0.3     goapp-consumer

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	configMap := &kafka.ConfigMap{
		"bootstrap.servers": "gokafka_kafka_1:9092",
		"client.id":         "goapp-consumer",
		"group.id":          "goapp-group",
		"auto.offset.reset": "earliest",
	}
	c, err := kafka.NewConsumer(configMap)
	if err != nil {
		fmt.Println("error consumer", err.Error())
	}
	topics := []string{"teste"}
	c.SubscribeTopics(topics, nil)
	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			fmt.Println(string(msg.Value), msg.TopicPartition)
		}
	}
}
