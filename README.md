# 📦 Order Processing System

## 📜 Overview

This project implements an Order Processing System composed of two microservices: **Order Service** and **Payment Service**. The services communicate asynchronously using **Kafka** as the message broker. 

- **Order Service**: Receives and processes orders.
- **Payment Service**: Handles payments when an order is placed.

![order processing](https://github.com/user-attachments/assets/9fe9407a-daa2-44c9-9d12-321673ee65ae)

## 🏗️ Architecture

The system is structured as follows:

1. **Order Service**: 
 - Sends an order creation event to Kafka whenever a new order is created.
 - Exposes a RESTful API to create new orders.

2. **Kafka**: 
 - Acts as the message broker that queues and distributes events between services.

3. **Payment Service**: 
 - Listens for order creation events from Kafka.
 - Processes the corresponding payments for the orders.


## 🚀 Getting Started

### 📋 Prerequisites

- Go (version 1.16 or higher)
- Kafka and Zookeeper running locally or accessible from your network.
- A database for each service (e.g., PostgreSQL, MySQL) configured.

### 💻 Installation

1. **Clone the repository**:
   ```bash
   git clone <your-repo-url>
   cd order-processing-system` 
	```
2.  **Install dependencies**: Make sure to run the necessary commands to install the required Go packages in both the Order Service and Payment Service directories.
    
3.  **Set up the database**: Ensure that the databases for both services are created and configured in the respective service configurations.
    
4.  **Start Kafka**: Make sure your Kafka and Zookeeper are up and running. You can use the following commands to start them if you have them installed locally:
    
   ``` bash
    # Start Zookeeper
    zookeeper-server-start.sh config/zookeeper.properties
    
    # Start Kafka
    kafka-server-start.sh config/server.properties` 
   ```

## 🔄 Running the Services

1.  **Start the Order Service**: Navigate to the Order Service directory and run:
    
    ```bash
    go run main.go` 
    ```
    This will start the Order Service, which listens on port 8080 by default.
    
2.  **Start the Payment Service**: Navigate to the Payment Service directory and run:
    
    ```bash
    go run main.go
    ```
    This will start the Payment Service, which listens on port 8081 by default.
    

## 🛍️ Creating an Order

To create an order, send a POST request to the Order Service:

```bash
`curl -X POST http://localhost:8080/orders \
     -H "Content-Type: application/json" \
     -d '{"id": 1, "amount": 100.0}'` 
```
## 💳 Payment Processing

Once the order is created, the Payment Service will listen for the `order.created` event on Kafka and process the payment accordingly.

## 📊 Sequence Diagram

Below is the sequence diagram illustrating the flow of events in the Order Processing System:
  

```sql
                        +-----------------+
                        |   Client Request |
                        +-----------------+
                                 |
                                 v
+-----------------------------+-------------------------------+
|                            Order Service                      |
|    +----------------------+   +---------------------------+   |
|    |     Database:       |   |      Message Queue        |    |
|    |      OrderDB        |   |    (Kafka Topic:          |    |
|    |                      |   |       order.created)      |   |
|    +----------------------+   +---------------------------+   |
|                |                          |                   |
|                | (1. Create Order)        |                   |
|                | (2. Persist Order)       |                   |
|                | (3. Publish Order Message)|                  |
|                v                          |                   |
|         +-----------------+                |                  |
|         |   Payment       |<---------------+                  |
|         |   Service       |                |                  |
|         +-----------------+                |                  |
|         |     Database:   |                |                  |
|         |   PaymentDB     |                |                  |
|         +-----------------+                |                  |
|                |                          |                   |
|                | (4. Process Payment)    |                    |
|                | (5. Publish Payment Result)|                 |
|                |                          |                   |
|                v                          |                   |
|         +-----------------+                |                  |
|         | Payment Gateway/ |                |                 |
|         |    Service        |                |                  |
|         +-----------------+                |                  |
|                |                          |                   |
|                | (6. Payment Result)      |                   |
|                |                          |                   |
|                v                          |                   |
|   +-----------------------------------------+                 |
|   |         Message Queue                   |                 |
|   |    (Kafka Topic: payment.completed      |                 |
|   |      or payment.failed)                 |                 |
|   +-----------------------------------------+                 |
+---------------------------------------------------------------+



```
## 📜 Logging

Both services log important events, including order creations and payment processing results. Check the console output for logs.

### ✅ To-Do List for Enhancing Order Processing System

1. **Authentication**
   - [ ] Implement user accounts and secure logins.
   
2. **Payment Integration**
   - [ ] Use a real payment gateway and handle sensitive data securely.
   
3. **Notifications**
   - [ ] Send confirmation emails and order updates to clients.
   
4. **Error Handling**
   - [ ] Improve user feedback and system reliability.
   
5. **Order Tracking**
   - [ ] Enable clients to track their orders.
   
6. **Testing**
   - [ ] Write tests to ensure the system is robust.



## 🤝 Contributing

If you'd like to contribute to this project, please fork the repository and submit a pull request.

## 📄 License

This project is licensed under the MIT License - see the LICENSE file for details.
