package app

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"order_service/config"

	httpservice "order_service/internal/adapter/http/service"
	postgresrepo "order_service/internal/adapter/postgres"
	"order_service/internal/usecase"
	"order_service/pkg/postgres"
)

const serviceName = "user-service"

type App struct {
	httpServer *httpservice.API
}

func New(ctx context.Context, cfg *config.Config) (*App, error) {
	log.Println(fmt.Sprintf("starting %v service", serviceName))

	log.Println("connecting to mongo", "database", cfg.Postgres.Dsn)
	postgresDB, err := postgres.New(ctx, cfg.Postgres)
	if err != nil {
		return nil, fmt.Errorf("mongo: %w", err)
	}

	orderRepo := postgresrepo.NewOrderRepository(postgresDB)

	// UseCase
	orderUsecase := usecase.NewOrder(orderRepo)

	// http service
	httpServer := httpservice.New(cfg.Server, orderUsecase)

	app := &App{
		httpServer: httpServer,
	}

	return app, nil
}

func (a *App) Close() {
	err := a.httpServer.Stop()
	if err != nil {
		log.Println("failed to shutdown gRPC service", err)
	}
}

func (a *App) Run() error {
	errCh := make(chan error, 1)

	a.httpServer.Run(errCh)

	log.Println(fmt.Sprintf("service %v started", serviceName))

	// Waiting signal
	shutdownCh := make(chan os.Signal, 1)
	signal.Notify(shutdownCh, syscall.SIGINT, syscall.SIGTERM)

	select {
	case errRun := <-errCh:
		return errRun

	case s := <-shutdownCh:
		log.Println(fmt.Sprintf("received signal: %v. Running graceful shutdown...", s))

		a.Close()
		log.Println("graceful shutdown completed!")
	}

	return nil
}
