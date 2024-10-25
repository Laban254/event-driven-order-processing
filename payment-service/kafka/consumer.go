package kafka

import (
    "encoding/json"
    "log"
    "payment-service/models" 
    "payment-service/services"
    "github.com/confluentinc/confluent-kafka-go/kafka"
)


func SetupPaymentConsumer(broker string) error {
    consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
        "bootstrap.servers": broker,
        "group.id":          "payment_service_group",
        "auto.offset.reset": "earliest",
    })
    if err != nil {
        return err
    }

    if err := consumer.Subscribe("order.created", nil); err != nil {
        return err
    }

    // Run the consumer in a goroutine
    go func() {
        for {
            msg, err := consumer.ReadMessage(-1) // Block until message is received
            if err == nil {
                var orderData map[string]interface{} // Use a map to hold dynamic data
                if err := json.Unmarshal(msg.Value, &orderData); err != nil {
                    log.Printf("Error unmarshalling message: %v", err)
                    continue
                }

                orderDataJSON, _ := json.Marshal(orderData) 
                log.Printf("Received order: %s\n", orderDataJSON)

                // Extract necessary fields
                orderID, ok := orderData["id"].(float64) 
                if !ok {
                    log.Printf("Invalid Order ID: cannot process payment")
                    continue
                }
                amount, ok := orderData["amount"].(float64)
                if !ok {
                    log.Printf("Invalid Amount: cannot process payment")
                    continue
                }

                // Create a payment based on the received order
                payment := &models.Payment{
                    OrderID: uint(orderID), 
                    Amount:  amount,
                    Status:  "pending", 
                }

                if err := services.ProcessPayment(payment); err != nil {
                    log.Printf("Error processing payment: %v", err)
                    //  todo publish a failure message
                } else {
                    // todo Publish success event 
                    log.Printf("Successfully processed payment for Order ID: %d", payment.OrderID)
                }
            } else {
                log.Printf("Error reading message: %v", err)
            }
        }
    }()
    return nil
}
