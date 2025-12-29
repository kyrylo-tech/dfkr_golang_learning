package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	file, err := os.OpenFile("/var/log/current_time.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening log file:", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(time.Now().Format("2006-01-02 15:04:05") + "\n")
	if err != nil {
		fmt.Println("Error writing to log file:", err)
	}
}
