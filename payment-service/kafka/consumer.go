package kafka

import (
    "payment-service/models"
    "payment-service/services"
    "log"

    "github.com/confluentinc/confluent-kafka-go/kafka"
)

var paymentProducer *kafka.Producer

// Change to a variable instead of a constant for the topic
var paymentProcessedTopic = "payment.processed"

// SetupPaymentProducer initializes the Kafka payment producer
func SetupPaymentProducer(broker string) {
    var err error
    paymentProducer, err = kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": broker})
    if err != nil {
        log.Fatalf("Failed to create payment producer: %s", err)
    }

    go func() {
        for e := range paymentProducer.Events() {
            switch ev := e.(type) {
            case *kafka.Message:
                if ev.TopicPartition.Error != nil {
                    log.Printf("Payment delivery failed: %v\n", ev.TopicPartition.Error)
                } else {
                    log.Printf("Payment message delivered to %v\n", ev.TopicPartition)
                }
            }
        }
    }()
}

// PublishPaymentProcessed publishes a payment processed event to Kafka
func PublishPaymentProcessed(payment models.Payment) error {
    // Call the CreatePayment service
    if err := services.CreatePayment(&payment); err != nil {
        log.Printf("Error creating payment: %v", err)
        return err // Handle error accordingly
    }

    jsonValue := payment.ToJSON() // Call the ToJSON method from the models package

    // Use the variable directly; no need for address-of
    message := &kafka.Message{
        TopicPartition: kafka.TopicPartition{
            Topic:     &paymentProcessedTopic, // Use the variable directly
            Partition: kafka.PartitionAny,
        },
        Value: []byte(jsonValue),
    }

    err := paymentProducer.Produce(message, nil) // Non-blocking produce
    if err != nil {
        log.Printf("Error producing payment message: %v", err)
        return err // Return the error for handling in the API
    }

    // Wait for message deliveries before shutting down
    paymentProducer.Flush(15 * 1000) // Wait up to 15 seconds for messages to be delivered

    return nil // Indicate success
}
