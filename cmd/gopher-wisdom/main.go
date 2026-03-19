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

	log.Println("Gopher Wisdom: The watcher has taken its post in the tunnels.")

	for {
		select {
		case <-ticker.C:
			log.Println("Gopher Wisdom: Still digging, still healthy. All tunnels are clear.")
		case <-ctx.Done():
			log.Println("Gopher Wisdom: The sun sets. Returning to the burrow gracefully...	")
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
			log.Fatalf("The tunnels have collapsed: %s\n", err)
		}
	}()

	log.Println("Gopher Wisdom is manifest at :8080. Seek and ye shall find.")

	<-ctx.Done()

	log.Println("A sign from above! Commencing the Great Hibernation...")
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	srv.Shutdown(shutdownCtx)

	sqlDB, err := cfg.DB.DB()
	if err == nil {
		log.Println("Sealing the archives. Rest well, little Gopher.")
		sqlDB.Close()
	}
}
