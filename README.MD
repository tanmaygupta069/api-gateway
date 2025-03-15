# Stock Trading Microservices Project

This project implements a stock trading platform using Go, gRPC, and Gin, following a microservices architecture.

## Architecture Overview
The system is designed with the following microservices:

- **Auth Service**: Handles user authentication and authorization (uses MySQL).
- **Order Service**: Manages stock orders and holdings (uses MySQL and Redis).
- **API Gateway**: Acts as the entry point for all client requests, routing them to the appropriate services.

### Architecture Diagram

![architecture diagram](https://raw.githubusercontent.com/tanmaygupta069/api-gateway/refs/heads/main/assets/stock-trading-architecture.png)


## Database Choices
- **MySQL**: ACID properties ensure consistency and durability which is crutial for a transactional use case as ours, Ideal for storing relational data:
  - User information (Auth Service).
  - Orders and Holdings data (Order Service).
- **Redis**: Chosen for caching in the Order Service to improve read performance for frequently accessed data like stock prices.

---

## Prerequisites
Ensure you have the following installed:
- **Docker** and **Docker Compose**
- **Go 1.22+**
- **MySQL**
- **Redis**

---

## Steps to Run the Project Locally
1. **Clone the Repositories**
   ```bash
   git clone https://github.com/tanmaygupta069/auth-service-go.git
   git clone https://github.com/tanmaygupta069/order-service-go.git
   git clone https://github.com/tanmaygupta069/api-gateway-go.git

## 2. Environment Variables
Set up `.env` files in each microservice with required database credentials.

### Example `.env` for Auth Service:
```env
JWT_SECRET = your-jwt-secret

MYSQL_PORT=3306
MYSQL_USER="root"
MYSQL_HOST="localhost"
MYSQL_DOCKER_HOST = "auth-mysql"
MYSQL_PASS= your-sql-pass
MYSQL_DB="traders"
MYSQL_ROOT_PASSWORD=your-sql-root-password

POST_SERVICE_PORT="8081"
POST_SERVICE_HOST = "localhost"
```

### Example `.env` for Order Service:
```env
PORT = "8080"

JWT_SECRET = your-jwt-secret
STOCK_API_KEY =  <!-- get api key from https://finnhub.io/dashboard -->

MYSQL_PORT=3306
MYSQL_USER="root"
MYSQL_HOST="localhost"
MYSQL_DOCKER_HOST = "order-mysql"
MYSQL_PASS= your-sql-pass
MYSQL_DB="traders"
MYSQL_ROOT_PASSWORD=your-sql-root-password

REDIS_PORT=6379
REDIS_USER=your-redis-user
REDIS_DB=0
REDIS_PASS=your-redis-pass
REDIS_HOST="localhost"

ORDER_SERVICE_PORT="8082"
ORDER_SERVICE_USER = "localhost"

BUCKET_SIZE = 10
RATE_LIMIT = 2
```

### Example `.env` for Api-Gateway:
```env
PORT="8080"

BUCKET_SIZE=10
RATE_LIMIT=2

AUTH_GRPC_PORT = "8081"
AUTH_GRPC_HOST = "localhost"

ORDER_GRPC_PORT = "8082"
ORDER_GRPC_HOST = "localhost"
```

## 3. Build and Start Containers
Use separate Docker Compose files to run each service:

### Start Auth Service:
Navigate to the `auth-service-go` directory and run:
```bash
docker-compose up --build
```

Repeat this for other directories as well `order-service-go` and `api-gateway-go`

## API Endpoints
Below is a brief description of the available API endpoints:

---

### **1. Signup**  
Registers a new user by providing a username and password.  

**Endpoint:** `POST /signup`  

---

### **2. Login**  
Authenticates a user and returns a JWT token for authorization.  

**Endpoint:** `POST /login`  

---

### **3. Place Order**  
Places a new stock order with stock ID, quantity, and price.After hitting this the orderStatus in the orders table will be set to "placed"

**Endpoint:** `POST /auth/order`  

---



### **4. Cancel Order**  
Cancels an existing order using the order ID. After hitting this the orderStatus in the orders table will be set to "cancelled"  

**Endpoint:** `PUT /auth/order`  

---


### **5. Get Order History**  
Retrieves the user's past order history by userId which is fetched from jwt token 

**Endpoint:** `GET /auth/orders`  

---



### **6. Get Current Price**  
Takes symbol and returns there current price from "https://finnhub.io" if symbol not present of incorrect returns a randomly assigned value which changes on each call to simulate realtime stock environment

**Endpoint:** `GET /stock/?symbol=""`  

---

### **7. Complete Order**  
Manually completes an order (used for testing purposes).Just a hack around to complete orders and updated them to the holdings

**Endpoint:** `POST /order/complete`  

---


### **8. Get Holdings**  
Returns the user's current stock holdings and total purchase price.  

**Endpoint:** `GET /order/holdings`


## Testing
For more intuitive testing i am attaching a postman collection as :
`api-gateway.postman_collection.json`
