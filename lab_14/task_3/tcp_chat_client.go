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

	reader := bufio.NewReader(os.Stdin)
	serverReader := bufio.NewReader(conn)

	fmt.Print("Enter nickname: ")
	name, _ := reader.ReadString('\n')
	conn.Write([]byte(name))

	go func() {
		for {
			msg, err := serverReader.ReadString('\n')
			if err != nil {
				os.Exit(0)
			}
			fmt.Print(msg)
		}
	}()

	fmt.Println("Connected. Type messages:")

	for {
		text, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		conn.Write([]byte(text))
	}
}
