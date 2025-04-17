package app

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"order_service/config"

	grpcserver "order_service/internal/adapter/grpc/server"
	"order_service/internal/adapter/http/myrouter"
	httpservice "order_service/internal/adapter/http/service"
	postgresrepo "order_service/internal/adapter/postgres"
	"order_service/internal/usecase"
	"order_service/pkg/postgres"
)

const serviceName = "Order"

type App struct {
	httpServer *httpservice.API
	postgresDB *postgres.PostgreDB
	grpcServer *grpcserver.API
}

func New(ctx context.Context, cfg *config.Config) (*App, error) {
	log.Printf("starting %v service\n", serviceName)

	log.Println("connecting to postgres")
	postgresDB, err := postgres.New(ctx, cfg.Postgres)
	if err != nil {
		return nil, fmt.Errorf("mongo: %w", err)
	}

	log.Println("connection established")

	// Repository
	orderRepo := postgresrepo.NewOrderRepository(postgresDB.Pool)

	// Inventory Service
	inv_router, err := myrouter.NewInventoryRouter("http://localhost:8082") // HardCode
	if err != nil {
		return nil, fmt.Errorf("inventory router: %w", err)
	}

	// inventoryServiceGRPCConn, err := grpcconn.New(cfg.GRPC.GRPCInventory.InventoryServiceURL)

	// UseCase
	orderUsecase := usecase.NewOrder(orderRepo, inv_router)

	// http service
	httpServer := httpservice.New(cfg.Server, orderUsecase)

	grpcServer := grpcserver.New(cfg.Server.GRPCServer, orderUsecase)

	app := &App{
		httpServer: httpServer,
		postgresDB: postgresDB,
		grpcServer: grpcServer,
	}

	return app, nil
}

// TODO: close postgres connection
func (a *App) Close(ctx context.Context) {
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

func (a *App) Run() error {
	errCh := make(chan error, 1)
	ctx := context.Background()

	// Running http server
	a.httpServer.Run(errCh)

	// Running gRPC server
	a.grpcServer.Run(ctx, errCh)

	log.Printf("service %v started\n", serviceName)

	// Waiting signal
	shutdownCh := make(chan os.Signal, 1)
	signal.Notify(shutdownCh, syscall.SIGINT, syscall.SIGTERM)

	select {
	case errRun := <-errCh:
		return errRun

	case s := <-shutdownCh:
		log.Printf("received signal: %v. Running graceful shutdown...\n", s)

		a.Close(ctx)
		log.Println("graceful shutdown completed!")
	}

	return nil
}
