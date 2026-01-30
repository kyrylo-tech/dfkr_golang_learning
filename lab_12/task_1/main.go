package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func main() {
	showNumbers := flag.Bool("n", false, "show line numbers")
	flag.Parse()

	files := flag.Args()

	if len(files) == 0 {
		fmt.Println("No files provided")
		return
	}

	for _, filename := range files {
		file, err := os.Open(filename)
		if err != nil {
			fmt.Println("Error opening file:", filename, err)
			continue
		}

		scanner := bufio.NewScanner(file)
		lineNum := 1

		for scanner.Scan() {
			if *showNumbers {
				fmt.Printf("%d %s\n", lineNum, scanner.Text())
			} else {
				fmt.Println(scanner.Text())
			}
			lineNum++
		}

		file.Close()
	}
}
