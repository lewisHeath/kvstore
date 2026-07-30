package main

import (
	"flag"

	"github.com/lewisHeath/kvstore/server"
)

func main() {
	s := server.NewServer()
	port := flag.String("port", "3000", "The port the server listens on")
	flag.Parse()
	s.Listen(*port)
}
