# event-driven-order-processing
event-driven-order-processing

 sudo  /usr/local/kafka/bin/kafka-server-start.sh  /usr/local/kafka/config/server.properties

zookeeper
 sudo /usr/local/kafka/bin/zookeeper-server-start.sh /usr/local/kafka/config/zookeeper.properties








```sql
                        +-----------------+
                        |   Client Request |
                        +-----------------+
                                 |
                                 v
+-----------------------------+-------------------------------+
|                            Order Service                      |
|    +----------------------+   +---------------------------+  |
|    |     Database:       |   |      Message Queue        |  |
|    |      OrderDB        |   |    (Kafka Topic:          |  |
|    |                      |   |       order.created)      |  |
|    +----------------------+   +---------------------------+  |
|                |                          |                  |
|                | (1. Create Order)        |                  |
|                | (2. Persist Order)       |                  |
|                | (3. Publish Order Message)|                  |
|                v                          |                  |
|         +-----------------+                |                  |
|         |   Payment       |<---------------+                  |
|         |   Service       |                |                  |
|         +-----------------+                |                  |
|         |     Database:   |                |                  |
|         |   PaymentDB     |                |                  |
|         +-----------------+                |                  |
|                |                          |                  |
|                | (4. Process Payment)    |                  |
|                | (5. Publish Payment Result)|                |
|                |                          |                  |
|                v                          |                  |
|         +-----------------+                |                  |
|         | Payment Gateway/ |                |                  |
|         |    Service        |                |                  |
|         +-----------------+                |                  |
|                |                          |                  |
|                | (6. Payment Result)      |                  |
|                |                          |                  |
|                v                          |                  |
|   +-----------------------------------------+                  |
|   |         Message Queue                   |                  |
|   |    (Kafka Topic: payment.completed      |                  |
|   |      or payment.failed)                 |                  |
|   +-----------------------------------------+                  |
+---------------------------------------------------------------+

```