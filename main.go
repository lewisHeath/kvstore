package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/lewisHeath/kvstore/server"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.Println("Application started....")

	s := server.NewServer()
	tcpPort := envOrDefault("TCP_PORT", "3000")
	httpPort := envOrDefault("HTTP_PORT", "8080")
	go s.Listen(tcpPort)
	go s.ServeHTTP(httpPort)

	<-ctx.Done()
	log.Println("Shutting down...")
}

func envOrDefault(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
