FROM golang:1.23

WORKDIR /app

COPY . . 

RUN go build -o inventory-service ./cmd/main.go

EXPOSE 8010

CMD ["./inventory-service"]