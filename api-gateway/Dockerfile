FROM golang:1.23

WORKDIR /app

COPY . . 

RUN go build -o api-gateway ./main.go

EXPOSE 8010

CMD ["./api-gateway"]