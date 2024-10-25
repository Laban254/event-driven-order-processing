# event-driven-order-processing
event-driven-order-processing

 sudo  /usr/local/kafka/bin/kafka-server-start.sh  /usr/local/kafka/config/server.properties

zookeeper
 sudo /usr/local/kafka/bin/zookeeper-server-start.sh /usr/local/kafka/config/zookeeper.properties















flow
Client Request
      |
      v
[order-service] -------> (1. Create Order)  
      |                   (2. Persist Order)  
      |                   (3. Publish to Kafka)  
      |  
      v  
[Kafka Topic: order.created] --------> (4. Triggered by new message)  
      |  
      v  
[payment-service] -------> (5. Process Payment)  
      |                   (6. Publish Payment Result)  
      |  
      v  
[Kafka Topic: payment.completed or payment.failed]  

```scss
          +-----------------+
          |  order-service  |
          +-----------------+
                   |
                   | (1. Publish order created message)
                   v
          [Kafka Topic: order.created]
                   |
                   v
          +------------------+
          |  payment-service  |
          +------------------+
                   |
                   | (2. Process payment)
                   v
          [Payment Gateway/Service]
                   |
                   | (3. Publish payment result)
                   v
          [Kafka Topic: payment.completed or payment.failed]
```

```sql
          +-----------------+
          |  order-service  |
          +-----------------+
          |  Database:      |
          |  OrderDB        |
          +-----------------+
                   |
                   | (1. Publish order created message)
                   v
          [Kafka Topic: order.created]
                   |
                   v
          +------------------+
          |  payment-service  |
          +------------------+
          |  Database:       |
          |  PaymentDB       |
          +------------------+
```