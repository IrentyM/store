FROM golang:1.24-slim

WORKDIR /app

COPY . . 

RUN go build -o inventory-service ./cmd/main.go

EXPOSE 8010

CMD ["./inventory-service"]