package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Config struct {
	LogLevel    string `json:"logLevel"`
	LogInterval int    `json:"logInterval"`
}

var cfg Config

func loadConfig() {
	f, err := os.ReadFile("config.json")
	if err != nil {
		fmt.Println("Error loading config:", err)
		return
	}

	err = json.Unmarshal(f, &cfg)
	if err != nil {
		fmt.Println("Error parsing config:", err)
		return
	}

	fmt.Println("Config reloaded:", cfg)
}

func main() {
	pid := os.Getpid()

	fmt.Printf("=== SIGUSR1 Config Reload Demo ===\n")
	fmt.Printf("Send reload with:\n  kill -USR1 %d\n\n", pid)

	loadConfig()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGUSR1)

	go func() {
		for s := range sigChan {
			fmt.Println("\n--- Received:", s, " ---")
			loadConfig()
		}
	}()

	for {
		fmt.Printf("[%s] Running... (interval %d sec)\n", cfg.LogLevel, cfg.LogInterval)
		time.Sleep(time.Duration(cfg.LogInterval) * time.Second)
	}
}
