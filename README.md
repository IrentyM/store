# Food Store

The `food-store` project is a microservices-based application for managing an e-commerce platform. It consists of the following services:

- **Inventory Service**: Manages product inventory and categories.
- **Order Service**: Handles customer orders and order items.
- **API Gateway**: Acts as a single entry point for routing requests to the appropriate services.

## Features

- **Inventory Service**:
  - Manage product categories (create, retrieve, update, delete, list).
  - Manage products (create, retrieve, update, delete, list).
  - Track stock and reserved quantities for products.

- **Order Service**:
  - Manage customer orders (create, retrieve, update, delete, list).
  - Manage order items associated with orders.

- **API Gateway**:
  - Routes requests to the appropriate microservices.
  - Provides a unified entry point for the application.

## Technologies Used

- **Programming Language**: Go (Golang)
- **Web Framework**: [Gin](https://github.com/gin-gonic/gin)
- **Database**: PostgreSQL
- **Containerization**: Docker, Docker Compose
- **Environment Management**: [godotenv](https://github.com/joho/godotenv)

## Project Structure

```
food-store/
├── .env                     # Environment variables
├── docker-compose.yml       # Docker Compose configuration
├── LICENSE                  # License file
├── README.md                # Project documentation
├── api-gateway/             # API Gateway service
│   ├── Dockerfile
│   ├── go.mod
│   ├── go.sum
│   └── main.go
├── db/                      # Database initialization scripts
│   └── init.sql
├── inventory-service/       # Inventory service
│   ├── Dockerfile
│   ├── go.mod
│   ├── go.sum
│   ├── cmd/
│   ├── internal/
│   ├── pkg/
│   └── README.md
├── order-service/           # Order service
│   ├── Dockerfile
│   ├── go.mod
│   ├── go.sum
│   ├── cmd/
│   ├── internal/
│   ├── pkg/
│   └── README.md
```

## Services

### Inventory Service

- **Base URL**: `http://localhost:8010`
- **Endpoints**:
  - `/categories`: Manage product categories.
  - `/products`: Manage products.
  - `/health`: Health check endpoint.

### Order Service

- **Base URL**: `http://localhost:8020`
- **Endpoints**:
  - `/orders`: Manage orders and order items.
  - `/health`: Health check endpoint.

### API Gateway

- **Base URL**: `http://localhost:8000`
- **Endpoints**:
  - `/inventory/*`: Routes requests to the inventory service.
  - `/orders/*`: Routes requests to the order service.
  - `/health`: Health check endpoint for the API Gateway.

## Running the Application

### Prerequisites

- Docker
- Docker Compose

### Steps

1. Clone the repository:
   ```sh
   git clone https://github.com/IrentyM/food-store.git
   cd food-store
   ```

2. Start the application using Docker Compose:
   ```sh
   docker-compose up --build
   ```

3. Access the services:
   - API Gateway: [http://localhost:8000](http://localhost:8000)
   - Inventory Service: [http://localhost:8010](http://localhost:8010)
   - Order Service: [http://localhost:8020](http://localhost:8020)

4. Verify the health of the services:
   - API Gateway: [http://localhost:8000/health](http://localhost:8000/health)
   - Inventory Service: [http://localhost:8010/health](http://localhost:8010/health)
   - Order Service: [http://localhost:8020/health](http://localhost:8020/health)

### Stopping the Application

To stop the application:
```sh
docker-compose down
```

### Cleaning Up

To remove all containers, networks, and volumes created by Docker Compose:
```sh
docker-compose down --volumes
```

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.