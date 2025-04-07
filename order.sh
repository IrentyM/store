#!/bin/bash

# Base directory for the order-service
BASE_DIR="order-service"

# Create the folder structure
mkdir -p $BASE_DIR/cmd
mkdir -p $BASE_DIR/internal/domain
mkdir -p $BASE_DIR/internal/usecase
mkdir -p $BASE_DIR/internal/repository
mkdir -p $BASE_DIR/internal/delivery/http
mkdir -p $BASE_DIR/internal/dto
mkdir -p $BASE_DIR/pkg/database
mkdir -p $BASE_DIR/pkg/logger
mkdir -p $BASE_DIR/configs
mkdir -p $BASE_DIR/migrations

# Create placeholder files
touch $BASE_DIR/cmd/main.go
touch $BASE_DIR/internal/domain/order.go
touch $BASE_DIR/internal/domain/order_item.go
touch $BASE_DIR/internal/usecase/order_usecase.go
touch $BASE_DIR/internal/usecase/order_usecase_test.go
touch $BASE_DIR/internal/repository/order_repository.go
touch $BASE_DIR/internal/repository/order_repository_postgres.go
touch $BASE_DIR/internal/repository/order_item_repository.go
touch $BASE_DIR/internal/repository/order_item_repository_postgres.go
touch $BASE_DIR/internal/delivery/http/handler.go
touch $BASE_DIR/internal/delivery/http/order_handler.go
touch $BASE_DIR/internal/delivery/http/order_item_handler.go
touch $BASE_DIR/internal/dto/order_dto.go
touch $BASE_DIR/internal/dto/order_item_dto.go
touch $BASE_DIR/pkg/database/postgres.go
touch $BASE_DIR/pkg/logger/logger.go
touch $BASE_DIR/configs/config.yaml
touch $BASE_DIR/migrations/001_init_orders_schema.up.sql
touch $BASE_DIR/go.mod
touch $BASE_DIR/go.sum
touch $BASE_DIR/README.md

echo "Folder structure for order-service created successfully!"