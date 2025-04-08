package service

import (
	"context"
	"errors"
	"fmt"
	"inventory_service/config"
	"inventory_service/internal/adapter/http/service/handler"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

const serverIPAddress = "127.0.0.1:%d" // Changed to 0.0.0.0 for external access

type API struct {
	server *gin.Engine
	cfg    config.HTTPServer
	addr   string

	inventoryHandler *handler.Inventory
}

func New(cfg config.Server, inventoryUseCase InventoryUsecase) *API {
	// Setting the Gin mode
	gin.SetMode(cfg.HTTPServer.Mode)
	// Creating a new Gin Engine
	server := gin.New()

	// Applying middleware
	server.Use(gin.Logger())
	server.Use(gin.Recovery())

	// Binding inventory
	inventoryHandler := handler.NewInventory(inventoryUseCase)

	api := &API{
		server:           server,
		cfg:              cfg.HTTPServer,
		addr:             fmt.Sprintf(serverIPAddress, cfg.HTTPServer.Port),
		inventoryHandler: inventoryHandler,
	}

	api.setupRoutes()

	return api
}

func (a *API) setupRoutes() {
	a.server.GET("/healthcheck", a.HealthCheck)

	products := a.server.Group("/products")
	{
		products.POST("/", a.inventoryHandler.Create)
		products.GET("/", a.inventoryHandler.GetList)
		products.GET("/:id", a.inventoryHandler.GetByID)
		products.PATCH("/:id", a.inventoryHandler.Update)
		products.DELETE("/:id", a.inventoryHandler.Delete)
	}
}

func (a *API) Stop() error {
	// Setting up the signal channel to catch termination signals
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// Blocking until a signal is received
	sig := <-quit
	log.Println("Shutdown signal received", "signal:", sig.String())

	// Creating a context with timeout for graceful shutdown
	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	log.Println("HTTP server shutting down gracefully")

	// Note: You can use `Shutdown` if you use `http.Server` instead of `gin.Engine`.
	log.Println("HTTP server stopped successfully")

	return nil
}

func (a *API) Run(errCh chan<- error) {
	go func() {
		log.Printf("HTTP server starting on: %v", a.addr)

		// No need to reinitialize `a.server` here. Just run it directly.
		if err := a.server.Run(a.addr); err != nil && !errors.Is(err, http.ErrServerClosed) {
			errCh <- fmt.Errorf("failed to start HTTP server: %w", err)
			return
		}
	}()
}

func (a *API) HealthCheck(c *gin.Context) {
	c.JSON(200, map[string]any{
		"status": "available",
		"system_info": map[string]string{
			"address": a.addr,
			"mode":    a.cfg.Mode,
		},
	})
}
