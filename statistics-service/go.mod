module github.com/IrentyM/store/statistics-service

go 1.23.5

require (
	github.com/IrentyM/store v0.0.0-00010101000000-000000000000
	github.com/caarlos0/env/v7 v7.1.0
	github.com/joho/godotenv v1.5.1
	github.com/lib/pq v1.10.9
	github.com/nats-io/nats.go v1.42.0
	google.golang.org/grpc v1.72.0
)

replace github.com/IrentyM/store => ../

require (
	github.com/klauspost/compress v1.18.0 // indirect
	github.com/nats-io/nkeys v0.4.11 // indirect
	github.com/nats-io/nuid v1.0.1 // indirect
	golang.org/x/crypto v0.37.0 // indirect
	golang.org/x/net v0.35.0 // indirect
	golang.org/x/sys v0.32.0 // indirect
	golang.org/x/text v0.24.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250218202821-56aae31c358a // indirect
	google.golang.org/protobuf v1.36.6 // indirect
)
