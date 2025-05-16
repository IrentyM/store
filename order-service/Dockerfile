FROM golang:1.23-slim


WORKDIR /app

COPY . . 

RUN go build -o order-service ./cmd/main.go

EXPOSE 8020

CMD ["./order-service"]