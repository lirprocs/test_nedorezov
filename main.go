package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"test_nedorezov/router"
	"time"
)

func main() {
	r := router.NewRouter()
	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	log.Printf("Starting server on %s...", server.Addr)

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)
	<-stop
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Failed to gracefully shutdown server: %v", err)
	}
	log.Println("Server stopped gracefully.")
}
