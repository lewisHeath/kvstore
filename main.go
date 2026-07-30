package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"net"
	"strings"
	"time"
)

var kv *Store

func main() {
	kv = NewStore()
	var port string
	flag.StringVar(&port, "port", "3000", "The port the server listens on")
	flag.Parse()
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
	buffer := make([]byte, 1024) // 1024 byte slice buffer to read the data into
	for {
		conn.SetReadDeadline(time.Now().Add(30 * time.Second))
		n, err := conn.Read(buffer)
		if err != nil {
			var netErr net.Error
			if errors.Is(err, io.EOF) {
				fmt.Printf("Client %v disconnected\n", conn.RemoteAddr())
			} else if errors.As(err, &netErr) && netErr.Timeout() {
				fmt.Printf("Client %v timed out\n", conn.RemoteAddr())
			} else {
				fmt.Printf("Error reading from the TCP connection %v %v\n", conn.RemoteAddr(), err)
			}
			return
		}
		// Output the data and echo back
		data := buffer[:n]
		s := strings.TrimRight(string(data), "\r\n")
		fmt.Printf("Received data from TCP socket: %s\n", s)
		// Handle the command
		r, err := handle(s, kv)
		if err != nil {
			fmt.Printf("Handle error from connection %v %v\n", conn.RemoteAddr(), err)
		}
		conn.Write([]byte(r + "\n"))
	}
}
