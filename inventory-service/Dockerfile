FROM golang:1.23

WORKDIR /app

COPY inventory-service/ . 

RUN go build -o inventory-service ./cmd/main.go

EXPOSE 8070

CMD ["./inventory-service"]