package main

import (
    "order-service/db"
    "order-service/kafka"
)

func main() {
    // Connect to the database
    db.ConnectDatabase()

    // Set up Kafka producer
    kafka.SetupProducer("localhost:9092") // Change this to your Kafka broker address

    // Keep the application running
    select {}
}
