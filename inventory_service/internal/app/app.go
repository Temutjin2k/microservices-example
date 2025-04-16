package app

import (
	"context"
	"fmt"
	"inventory_service/config"
	grpcserver "inventory_service/internal/adapter/grpc/server"
	httpservice "inventory_service/internal/adapter/http/service"
	postgresrepo "inventory_service/internal/adapter/postgres"
	"inventory_service/internal/usecase"
	"inventory_service/pkg/postgres"
	"log"
	"os"
	"os/signal"
	"syscall"
)

const serviceName = "Inventory"

type Application struct {
	httpServer *httpservice.API
	grpcServer *grpcserver.API
	postgresDB *postgres.PostgreDB
}

func New(ctx context.Context, config *config.Config) (*Application, error) {
	log.Printf("starting %v service\n", serviceName)
	log.Println("connecting to postgres")

	postgresDB, err := postgres.New(ctx, config.Postgres)
	if err != nil {
		return nil, fmt.Errorf("mongo: %w", err)
	}
	log.Println("connection established")

	inventoryRepo := postgresrepo.NewInventoryRepository(postgresDB.Pool)

	inventoryUseCase := usecase.NewInventory(inventoryRepo)

	httpServer := httpservice.New(config.Server, inventoryUseCase)

	grpcServer := grpcserver.New(config.Server.GRPCServer, inventoryUseCase)

	app := &Application{
		httpServer: httpServer,
		grpcServer: grpcServer,
		postgresDB: postgresDB,
	}

	return app, nil
}

func (a *Application) Close(ctx context.Context) {
	// Closing http server
	err := a.httpServer.Stop()
	if err != nil {
		log.Println("failed to shutdown service", err)
	}

	err = a.grpcServer.Stop(ctx)
	if err != nil {
		log.Println("failed to shutdown service", err)
	}

	// Closing postgres connection
	a.postgresDB.Pool.Close()
}

func (app *Application) Run() error {
	errCh := make(chan error, 1)
	ctx := context.Background()

	// Running http server
	app.httpServer.Run(errCh)

	// Running gRPC server
	app.grpcServer.Run(ctx, errCh)

	log.Printf("service %v started\n", serviceName)

	// Waiting signal
	shutdownCh := make(chan os.Signal, 1)
	signal.Notify(shutdownCh, syscall.SIGINT, syscall.SIGTERM)

	select {
	case errRun := <-errCh:
		return errRun

	case s := <-shutdownCh:
		log.Printf("received signal: %v. Running graceful shutdown...\n", s)

		app.Close(ctx)
		log.Println("graceful shutdown completed!")
	}

	return nil
}
