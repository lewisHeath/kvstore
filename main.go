package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"net"
	"strings"
	"time"

	"github.com/lewisHeath/kvstore/store"
)

var kv *store.Store

func main() {
	kv = store.NewStore()
	var port string
	flag.StringVar(&port, "port", "3000", "The port the server listens on")
	flag.Parse()

	listen(port)
}

func listen(port string) {
	fmt.Printf("Initialising TCP server on port %v\n", port)
	ln, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		fmt.Printf("Error starting the TCP server: %v\n", err)
		return
	}
	defer ln.Close()
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Printf("Error accepting a connection on listener %v port %v: %v\n", ln.Addr().Network(), ln.Addr().String(), err)
			return
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	fmt.Printf("Accepted TCP connection from %v\n", conn.RemoteAddr())
	scanner := bufio.NewScanner(conn)
	for {
		conn.SetReadDeadline(time.Now().Add(30 * time.Second))
		if !scanner.Scan() {
			if err := scanner.Err(); err != nil {
				var netErr net.Error
				if errors.As(err, &netErr) && netErr.Timeout() {
					fmt.Printf("Client %v timed out\n", conn.RemoteAddr())
				} else {
					fmt.Printf("Error reading from the TCP connection %v %v\n", conn.RemoteAddr(), err)
				}
				return
			} else {
				fmt.Printf("Client %v disconnected\n", conn.RemoteAddr())
			}
			return
		}
		line := scanner.Text()
		fmt.Printf("Received data from TCP socket: %s\n", line)
		r := handle(line, kv)        // Handle the command
		conn.Write([]byte(r + "\n")) // Add newline
	}
}

func handle(c string, kv *store.Store) string {
	s := strings.SplitN(c, " ", 3)
	if len(s) == 0 {
		return "EMPTY COMMAND"
	}
	switch s[0] {
	case "GET":
		if len(s) != 2 {
			return "GET requires 1 key"
		}
		fmt.Printf("Performing GET on key=%v\n", s[1])
		v, ok := kv.Get(s[1])
		if !ok {
			return "NOTFOUND"
		}
		return v
	case "PUT":
		if len(s) != 3 {
			return "PUT requires a 2 inputs, a key and a value"
		}
		fmt.Printf("Performing PUT with key=%v value=%v\n", s[1], s[2])
		kv.Put(s[1], s[2])
		return "OK"
	case "DELETE":
		if len(s) != 2 {
			return "DELETE requires 1 key"
		}
		fmt.Printf("Performing DELETE on key=%v\n", s[1])
		ok := kv.Delete(s[1])
		if !ok {
			return "NOTFOUND"
		}
		return "OK"
	default:
		return "UNRECOGNIZED COMMAND"
	}
}
