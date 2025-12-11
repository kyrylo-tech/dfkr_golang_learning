package main

import (
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"syscall"
)

func main() {
	pid := os.Getegid()
	fmt.Println("=== SIGQUIT Handler Demo ===")
	fmt.Printf("Send SIGQUIT with:\n  kill -3 %d\n", pid)
	fmt.Println("Or press: Ctrl+\n")

	sigChan := make(chan os.Signal, 1)

	signal.Notify(sigChan, syscall.SIGQUIT)

	go func() {
		for s := range sigChan {
			fmt.Println("\n--- Received signal:", s, "---")

			stack := make([]byte, 8192)
			n := runtime.Stack(stack, true)

			fmt.Println("=== STACK DUMP ===")
			fmt.Println(string(stack[:n]))
			fmt.Println("===================")
		}
	}()

	select {}
}
