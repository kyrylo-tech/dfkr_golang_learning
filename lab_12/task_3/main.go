package main

import (
	"bufio"
	"io"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) > 0 {
		return
	}

	reader := bufio.NewReader(os.Stdin)

	for {
		line, err := reader.ReadString('\n')

		if err == io.EOF {
			if len(line) > 0 {
				os.Stdout.WriteString(line)
			}
			break
		}

		if err != nil {
			break
		}

		os.Stdout.WriteString(line)
	}
}
