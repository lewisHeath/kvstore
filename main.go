package main

import (
	"context"
	"flag"
	"fmt"
	"os/signal"
	"syscall"

	"github.com/lewisHeath/kvstore/server"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	fmt.Println("Application started. Press Ctrl+C to exit.")

	s := server.NewServer()
	tcp_port := flag.String("tcp-port", "3000", "The tcp port the server listens on")
	http_port := flag.String("http-port", "8080", "The http port the server listens on")
	flag.Parse()
	go s.Listen(*tcp_port)
	go s.ServeHTTP(*http_port)

	<-ctx.Done()
	fmt.Println("Shutting down...")
}
