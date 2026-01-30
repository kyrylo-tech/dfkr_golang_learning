package main

import (
	"bufio"
	"fmt"
	"net"
)

func handleClient(conn net.Conn) {
	defer conn.Close()

	addr := conn.RemoteAddr().String()
	fmt.Println("Client connected:", addr)

	reader := bufio.NewReader(conn)

	for {
		msg, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Client disconnected:", addr)
			return
		}

		fmt.Printf("Received from %s: %s", addr, msg)

		conn.Write([]byte(msg))
	}
}

func main() {
	port := ":8083"

	listener, err := net.Listen("tcp", port)
	if err != nil {
		panic(err)
	}

	fmt.Println("TCP Echo server listening on", port)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		go handleClient(conn)
	}
}
