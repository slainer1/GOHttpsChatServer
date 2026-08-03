package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handleConnection(conn)
	}

}
func handleConnection(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	line, err := reader.ReadString('\n')
	if err != nil {
		fmt.Fprintf(conn, "Error reading command: %v\n", err)
		return
	}
	parts := strings.SplitN(strings.TrimSpace(line), " ", 2)
	if len(parts) != 2 {
		fmt.Fprintf(conn, "Error parsing command: %v\n", line)
		return
	}
	command := parts[0]
	response := parts[1]
	log.Printf("Received command: %s %s\n", command, response)

	switch command {
	case "GET":
		handleGet(conn, response)
	default:
		fmt.Fprintf(conn, "Unknown command: %s\n", command)
	}
}
func handleGet(conn net.Conn, response string) {
	fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
}
