package kafka

import (
    "order-service/models"
    "log"

    "github.com/confluentinc/confluent-kafka-go/kafka"
)

var (
    producer          *kafka.Producer
    orderCreatedTopic = "order.created"
)

func SetupProducer(broker string) {
    var err error
    producer, err = kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": broker})
    if err != nil {
        log.Fatalf("Failed to create producer: %s", err)
    }

    // Check for broker availability with metadata query
    metadata, err := producer.GetMetadata(nil, true, 5000) // Timeout of 5 seconds
    if err != nil {
        log.Fatalf("Failed to connect to Kafka broker: %s", err)
    } else {
        log.Printf("Connected to Kafka broker(s): %v", metadata.Brokers)
    }

    // Handle message delivery events
    go func() {
        for e := range producer.Events() {
            switch ev := e.(type) {
            case *kafka.Message:
                if ev.TopicPartition.Error != nil {
                    log.Printf("Delivery failed: %v\n", ev.TopicPartition.Error)
                } else {
                    log.Printf("Message delivered to %v\n", ev.TopicPartition)
                }
            }
        }
    }()
}

func PublishOrderCreated(order models.Order) error {

    jsonValue := order.ToJSON()
    message := &kafka.Message{
        TopicPartition: kafka.TopicPartition{Topic: &orderCreatedTopic, Partition: kafka.PartitionAny},
        Value:          []byte(jsonValue),
    }

    err := producer.Produce(message, nil) // Non-blocking produce
    if err != nil {
        log.Printf("Error producing message: %v", err)
        return err 
    }

    // Wait for message deliveries before shutting down
    producer.Flush(15 * 1000) // Wait up to 15 seconds for messages to be delivered
    log.Println("All messages flushed to Kafka successfully.")

    return nil // Indicate success
}
