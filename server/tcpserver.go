package server

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"net"
	"strings"
	"time"

	"github.com/lewisHeath/kvstore/store"
)

type Server struct {
	store *store.Store
}

func NewServer() *Server {
	return &Server{
		store: store.NewStore(),
	}
}

func (s *Server) Listen(port string) {
	log.Printf("Initialising TCP server on port %v\n", port)
	ln, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		fmt.Printf("Error starting the TCP server: %v\n", err)
		return
	}
	defer ln.Close()
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("Error accepting a connection on listener %v port %v: %v\n", ln.Addr().Network(), ln.Addr().String(), err)
			return
		}
		go s.handleConnection(conn)
	}
}

func (s *Server) handleConnection(conn net.Conn) {
	defer conn.Close()
	log.Printf("Accepted TCP connection from %v\n", conn.RemoteAddr())
	scanner := bufio.NewScanner(conn)
	conn.SetReadDeadline(time.Now().Add(30 * time.Second))
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		log.Printf("Received data from TCP socket: %s\n", line)
		r := s.handle(line)   // Handle the command
		fmt.Fprintln(conn, r) // Output to the connection
		conn.SetReadDeadline(time.Now().Add(30 * time.Second))
	}
	if err := scanner.Err(); err != nil {
		var netErr net.Error
		if errors.As(err, &netErr) && netErr.Timeout() {
			log.Printf("Client %v timed out\n", conn.RemoteAddr())
		} else {
			log.Printf("Error reading from the TCP connection %v %v\n", conn.RemoteAddr(), err)
		}
		return
	}
	log.Printf("client %v disconnected\n", conn.RemoteAddr())
}

func (s *Server) handle(c string) string {
	command := strings.SplitN(c, " ", 3)
	if len(command) == 0 {
		return "EMPTY COMMAND"
	}
	switch command[0] {
	case "GET":
		if len(command) != 2 {
			return "GET requires 1 key"
		}
		log.Printf("Performing GET on key=%v\n", command[1])
		v, ok := s.store.Get(command[1])
		if !ok {
			return "NOTFOUND"
		}
		return v
	case "PUT":
		if len(command) != 3 {
			return "PUT requires a 2 inputs, a key and a value"
		}
		log.Printf("Performing PUT with key=%v value=%v\n", command[1], command[2])
		s.store.Put(command[1], command[2])
		return "OK"
	case "DELETE":
		if len(command) != 2 {
			return "DELETE requires 1 key"
		}
		log.Printf("Performing DELETE on key=%v\n", command[1])
		ok := s.store.Delete(command[1])
		if !ok {
			return "NOTFOUND"
		}
		return "OK"
	default:
		return "UNRECOGNIZED COMMAND"
	}
}
