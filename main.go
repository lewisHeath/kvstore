package main

import (
	"context"
	"flag"
	"log"
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
	tcp_port := flag.String("tcp-port", "3000", "The tcp port the server listens on")
	http_port := flag.String("http-port", "8080", "The http port the server listens on")
	flag.Parse()
	go s.Listen(*tcp_port)
	go s.ServeHTTP(*http_port)

	<-ctx.Done()
	log.Println("Shutting down...")
}
