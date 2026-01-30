package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Fprintln(os.Stderr, "No files provided")
		return
	}

	for _, name := range args {
		file, err := os.Open(name)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error:", err)
			continue
		}

		_, err = io.Copy(os.Stdout, file)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Copy error:", err)
		}

		file.Close()
	}
}
