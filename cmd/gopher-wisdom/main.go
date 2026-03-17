package main

import (
	"context"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/EternalHalve/gopher-wisdom/internal/config"
	"github.com/EternalHalve/gopher-wisdom/internal/quotes"
	"github.com/gin-gonic/gin"
)

func startStatusWorker(ctx context.Context) {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	log.Println("Background: Worker started")

	for {
		select {
		case <-ticker.C:
			log.Println("Background: Server is healthy...")
		case <-ctx.Done():
			log.Println("Background: Worker shutting down gracefully...")
			return
		}
	}
}

func main() {
	cfg := config.Load()
	cfg.DB.AutoMigrate(&quotes.Quote{})

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	go startStatusWorker(ctx)

	router := gin.Default()
	quotes.RegisterRoutes(router, cfg.DB)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Listen error: %s\n", err)
		}
	}()

	log.Println("Server started on :8080")

	<-ctx.Done()

	log.Println("Shutting down...")
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	srv.Shutdown(shutdownCtx)
}
