FROM golang:1.24-slim

WORKDIR /app

COPY . . 

RUN go build -o api-gateway ./main.go

EXPOSE 8010

CMD ["./api-gateway"]