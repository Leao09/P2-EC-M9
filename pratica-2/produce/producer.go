package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func SensorData() map[string]interface{} {
    data := map[string]interface{}{
        "idSensor": "sensor_001",
		"Timestamp": time.Now(),
		"tipoPoluente": "PM2.5",
        "nivel":  rand.Intn(100),
    }
    return data
}

func Client() {
	producer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:29092,localhost:39092",
		"client.id":         "go-producer",
	})
	if err != nil {
		panic(err)
	}
	defer producer.Close()

	// Enviar mensagem
	topic := "test_topic"
	message := SensorData()
	messageBytes, err := json.Marshal(message)
	if err != nil {
		panic(err)
	}
	producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          messageBytes,
	}, nil) 

	for { 
		producer.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          messageBytes,
		}, nil) 
		log.Println("[Producer]] ", message)
		time.Sleep(2 * time.Second)
	}
}

func main() {
	Client()
}