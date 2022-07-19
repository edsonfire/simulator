package main

import (
	"fmt"
	kafka2 "https://github.com/edsonfire/simulator/tree/main/application/kafka"
	"https://github.com/edsonfire/simulator/tree/main/infra/kafka"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/joho/godotenv"
	"log"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
}

func main() {

producer:= kafka.newKafkaProducer()
kafka.Publish("ola", "readtest", producer)

	/*
	msgChan := make(chan *ckafka.Message)
	consumer := kafka.NewKafkaConsumer(msgChan)
	go consumer.Consume()
	for msg := range msgChan {
		fmt.Println(string(msg.Value))
		go kafka2.Produce(msg)
	}*/
}
