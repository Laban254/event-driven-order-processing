package kafka

import (
    "order-service/models"
    "order-service/services"
    "log"

    "github.com/confluentinc/confluent-kafka-go/kafka"
)

var (
    producer          *kafka.Producer
    orderCreatedTopic = "order.created" // Change to a variable
)

// SetupProducer initializes the Kafka producer
func SetupProducer(broker string) {
    var err error
    producer, err = kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": broker})
    if err != nil {
        log.Fatalf("Failed to create producer: %s", err)
    }

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

// PublishOrderCreated publishes an order created event to Kafka
func PublishOrderCreated(order models.Order) error {
    // Call the CreateOrder service
    if err := services.CreateOrder(&order); err != nil {
        log.Printf("Error creating order: %v", err)
        return err // Handle error accordingly
    }

    jsonValue := order.ToJSON() // Call the ToJSON method from the models package
    message := &kafka.Message{
        TopicPartition: kafka.TopicPartition{Topic: &orderCreatedTopic, Partition: kafka.PartitionAny},
        Value:          []byte(jsonValue),
    }

    err := producer.Produce(message, nil) // Non-blocking produce
    if err != nil {
        log.Printf("Error producing message: %v", err)
        return err // Return the error for handling in the API
    }

    // Wait for message deliveries before shutting down
    producer.Flush(15 * 1000) // Wait up to 15 seconds for messages to be delivered

    return nil // Indicate success
}
