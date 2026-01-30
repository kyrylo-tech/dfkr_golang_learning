package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8083")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	fmt.Println("Connected to TCP Echo server (localhost:8083)")
	fmt.Println("Enter text and press Enter:")

	serverReader := bufio.NewReader(conn)
	inputReader := bufio.NewReader(os.Stdin)

	for {
		text, err := inputReader.ReadString('\n')
		if err != nil {
			break
		}

		start := time.Now()

		_, err = conn.Write([]byte(text))
		if err != nil {
			fmt.Println("Send error:", err)
			continue
		}

		reply, err := serverReader.ReadString('\n')
		if err != nil {
			fmt.Println("Read error:", err)
			break
		}

		rtt := time.Since(start).Milliseconds()

		if strings.TrimSpace(text) == strings.TrimSpace(reply) {
			fmt.Printf("✔ Delivered. RTT = %d ms\n", rtt)
		} else {
			fmt.Printf("✖ Data mismatch. RTT = %d ms\n", rtt)
			fmt.Println("Sent:", text)
			fmt.Println("Got :", reply)
		}
	}
}
