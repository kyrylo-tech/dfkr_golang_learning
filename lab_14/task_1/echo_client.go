package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8083")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	fmt.Println("Connected to server. Type text:")

	serverReader := bufio.NewReader(conn)
	inputReader := bufio.NewReader(os.Stdin)

	for {
		text, _ := inputReader.ReadString('\n')
		conn.Write([]byte(text))

		reply, _ := serverReader.ReadString('\n')
		fmt.Print("Echo:", reply)
	}
}
