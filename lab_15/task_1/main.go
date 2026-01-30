package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

const pageSize = 10

func main() {
	showNumbers := flag.Bool("N", false, "show line numbers")
	flag.Parse()

	args := flag.Args()

	filename := ""
	for _, a := range args {
		if !strings.HasPrefix(a, "-") {
			filename = a
			break
		}
	}

	if filename == "" {
		fmt.Println("Missing filename.")
		return
	}

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	lineNumber := 1
	printed := 0

	for scanner.Scan() {
		line := scanner.Text()

		if *showNumbers {
			fmt.Printf("%d %s\n", lineNumber, line)
		} else {
			fmt.Println(line)
		}

		lineNumber++
		printed++

		if printed >= pageSize {
			break
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("\n%s\n", filename)
}
