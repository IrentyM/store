package app

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/IrentyM/store/statistics-service/config"
	grpcserver "github.com/IrentyM/store/statistics-service/internal/adapter/controller/grpc/server"
	natsadapter "github.com/IrentyM/store/statistics-service/internal/adapter/nats/handler"
	statisticsrepo "github.com/IrentyM/store/statistics-service/internal/adapter/repo/postgres"
	"github.com/IrentyM/store/statistics-service/internal/usecase"
	postgresconn "github.com/IrentyM/store/statistics-service/pkg/postgres"
	"github.com/nats-io/nats.go"
)

const serviceName = "statistics-service"

type App struct {
	grpcServer  *grpcserver.Server
	natsHandler *natsadapter.NATSHandler
	natsConn    *nats.Conn
}

func New(ctx context.Context, cfg *config.Config) (*App, error) {
	log.Printf("starting %v service", serviceName)

	// Connect to PostgreSQL
	log.Println("Connecting to PostgreSQL...")
	db, err := postgresconn.NewPostgreConnection(cfg.Postgres)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to PostgreSQL: %w", err)
	}
	log.Println("Connected to PostgreSQL successfully.")

	log.Println("connecting to NATS", "url", cfg.NATS.URL)
	nc, err := nats.Connect(cfg.NATS.URL)
	if err != nil {
		return nil, fmt.Errorf("nats connection failed: %w", err)
	}

	repo := statisticsrepo.NewStatsRepository(db)
	uc := usecase.NewStatsUseCase(repo)

	grpcServer, err := grpcserver.NewServer(cfg.Server.GRPCServer.Port, uc)
	if err != nil {
		return nil, fmt.Errorf("failed to create gRPC server: %w", err)
	}

	natsHandler := natsadapter.NewNATSHandler(nc, uc)

	return &App{
		grpcServer:  grpcServer,
		natsHandler: natsHandler,
		natsConn:    nc,
	}, nil
}

func (a *App) Close() {
	a.grpcServer.Stop()
	a.natsConn.Close()
}

func (a *App) Run() error {
	errCh := make(chan error, 1)

	go a.natsHandler.Start()

	go func() {
		if err := a.grpcServer.Start(); err != nil {
			log.Fatalf("gRPC server failed: %v", err)
		}
	}()

	log.Printf("service %v started", serviceName)

	shutdownCh := make(chan os.Signal, 1)
	signal.Notify(shutdownCh, syscall.SIGINT, syscall.SIGTERM)

	select {
	case errRun := <-errCh:
		return errRun

	case s := <-shutdownCh:
		log.Printf("received signal: %v. Running graceful shutdown...", s)
		a.Close()
		log.Println("graceful shutdown completed!")
	}

	return nil
}
