package main

import (
	"flag"
	"fmt"
	"io/fs"
	"path/filepath"
	"strings"
)

func main() {

	myPath := flag.String("path", ".", "Use directory to show tree")
	dotFiles := flag.Bool("dotfiles", false, "Skip dotfiles")
	enableStructure := flag.Bool("structure", false, "Enable structured tree view")

	flag.Parse()

	fmt.Printf("(kTree) Path: %s\n", *myPath)

	err := filepath.WalkDir(*myPath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if *dotFiles && path != *myPath && strings.HasPrefix(d.Name(), ".") {
			if d.IsDir() {
				return fs.SkipDir
			}
			return nil
		}

		if *enableStructure {
			relPath, _ := filepath.Rel(*myPath, path)
			if relPath == "." {
				fmt.Println(relPath)
				return nil
			}

			depth := strings.Count(relPath, string(filepath.Separator))

			indentation := strings.Repeat("    ", depth-1)
			prefix := "├── "
			fmt.Printf("%s%s%s\n", indentation, prefix, d.Name())

		} else {
			fmt.Println(path)
		}

		return nil
	})

	if err != nil {
		fmt.Printf("Error walking the path %q: %v\n", *myPath, err)
		return
	}
}
