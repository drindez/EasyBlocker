package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	ln, err := net.Listen("tcp", "0.0.0.0:0")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	fmt.Println("Listening on", ln.Addr())

	file, err := os.Create("connections.txt")
	if err != nil {
		log.Fatalf("Failed to create file: %v", err)
	}
	defer file.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("Failed to accept connection: %v", err)
			continue
		}

		remoteAddr := conn.RemoteAddr().(*net.TCPAddr)
		ip := remoteAddr.IP.String()

		fmt.Println("New connection from:", ip)

		if _, err := file.WriteString(ip + "\n"); err != nil {
			log.Printf("Failed to write IP address to file: %v", err)
		}

		conn.Close()
	}
}
