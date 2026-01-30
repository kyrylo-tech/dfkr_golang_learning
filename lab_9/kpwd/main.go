package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {

	logical := flag.Bool("L", false, "Вивести логічний шлях")
	physical := flag.Bool("P", false, "Вивести фізичний шлях")

	help := flag.Bool("help", false, "Показати допомогу")
	version := flag.Bool("version", false, "Show version")
	author := flag.Bool("author", false, "Show author")

	flag.Parse()

	if *help {
		fmt.Println("Аналог утиліти pwd (Go):")
		fmt.Println(" -L (default) показати логічний шлях")
		fmt.Println(" -P показати фізичний шлях, без символічних лінків")
		fmt.Println(" --help показати довідку")
		fmt.Println(" --version показати версію утиліти")
		fmt.Println(" --author показати ім’я та прізвище розробника")
		os.Exit(0)
	} else if *version {
		fmt.Println("Version: v.01")
	} else if *author {
		fmt.Println("Author of this shit:")
		fmt.Println("Kyrylo Bitsay, 74 group")
	} else {
		if !*logical && !*physical {
			*logical = true
		}

		if *logical {
			pwd := os.Getenv("PWD")
			if pwd == "" {
				pwd, _ = os.Getwd()
			}
			fmt.Println(pwd)
		} else if *physical {
			dir, err := os.Getwd()
			if err != nil {
				log.Fatal("Не вдалося отримати фізичний шлях:", err)
			}
			realPath, err := filepath.EvalSymlinks(dir)
			if err != nil {
				log.Fatal("Не вдалося вирішити символічні лінки:", err)
			}
			fmt.Println(realPath)
		}
	}

}
