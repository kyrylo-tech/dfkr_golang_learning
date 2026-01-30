package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

type Client struct {
	name string
	conn net.Conn
	ch   chan string
}

var (
	joinCh  = make(chan Client)
	leaveCh = make(chan Client)
	msgCh   = make(chan string)
	clients = make(map[net.Conn]Client)
	history []string
	maxHist = 100
)

func broadcaster() {
	for {
		select {

		case c := <-joinCh:
			clients[c.conn] = c
			sys := fmt.Sprintf("[SYSTEM] %s joined\n", c.name)
			addHistory(sys)
			broadcast(sys)

			// надіслати історію новому клієнту
			for _, h := range history {
				c.ch <- h
			}

		case c := <-leaveCh:
			delete(clients, c.conn)
			close(c.ch)
			sys := fmt.Sprintf("[SYSTEM] %s left\n", c.name)
			addHistory(sys)
			broadcast(sys)

		case msg := <-msgCh:
			addHistory(msg)
			broadcast(msg)
		}
	}
}

func broadcast(msg string) {
	for _, c := range clients {
		select {
		case c.ch <- msg:
		default:
		}
	}
}

func addHistory(msg string) {
	history = append(history, msg)
	if len(history) > maxHist {
		history = history[len(history)-maxHist:]
	}
}

func handleClient(conn net.Conn) {
	reader := bufio.NewReader(conn)

	nameRaw, err := reader.ReadString('\n')
	if err != nil {
		conn.Close()
		return
	}
	name := strings.TrimSpace(nameRaw)

	client := Client{
		name: name,
		conn: conn,
		ch:   make(chan string, 10),
	}

	joinCh <- client

	go func() {
		for msg := range client.ch {
			conn.Write([]byte(msg))
		}
	}()

	for {
		text, err := reader.ReadString('\n')
		if err != nil {
			break
		}

		text = strings.TrimSpace(text)
		if text == "" {
			continue
		}

		full := fmt.Sprintf("[%s]: %s\n", name, text)
		msgCh <- full
	}

	leaveCh <- client
	conn.Close()
}

func main() {
	listener, err := net.Listen("tcp", ":8083")
	if err != nil {
		panic(err)
	}

	fmt.Println("TCP Chat server running on :8083")

	go broadcaster()

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go handleClient(conn)
	}
}
