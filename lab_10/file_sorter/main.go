package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func fileCategory(path string) string {
	ext := strings.ToLower(filepath.Ext(path))
	switch ext {
	case ".txt", ".md":
		return "Documents"
	case ".go":
		return "Sources"
	case ".jpg", ".jpeg", ".png":
		return "Images"
	case ".pdf":
		return "PDFs"
	default:
		return "Other"
	}
}

func moveFile(src, destDir string, logger *log.Logger) error {
	if err := os.MkdirAll(destDir, 0o755); err != nil {
		logger.Printf("ERROR creating dir %s: %v\n", destDir, err)
		return err
	}

	destPath := filepath.Join(destDir, filepath.Base(src))

	if err := os.Rename(src, destPath); err != nil {
		logger.Printf("ERROR moving %s -> %s: %v\n", src, destPath, err)
		return err
	}

	msg := fmt.Sprintf("Переміщено: %s -> %s\n", src, destPath)
	fmt.Print(msg)
	logger.Print(msg)

	return nil
}

func main() {
	source := flag.String("source", "", "Source directory")
	dest := flag.String("dest", "", "Destination directory")
	flag.Parse()

	if *source == "" || *dest == "" {
		fmt.Println("Використання: go run file_sorter.go --source=SRC --dest=DEST")
		return
	}

	if err := os.MkdirAll(*dest, 0o755); err != nil {
		log.Fatal("Не вдалося створити dest:", err)
	}

	logFilePath := filepath.Join(*dest, "sort_report.log")
	logFile, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o644)
	if err != nil {
		log.Fatal("Не вдалося відкрити log файл:", err)
	}
	defer logFile.Close()

	logger := log.New(io.MultiWriter(logFile), "", log.LstdFlags)

	fmt.Println("Програма сортування запущена...")
	logger.Println("=== Start sorting ===")

	sourceAbs, _ := filepath.Abs(*source)
	destAbs, _ := filepath.Abs(*dest)

	err = filepath.WalkDir(sourceAbs, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			logger.Printf("ERROR accessing %s: %v\n", path, err)
			return nil
		}

		if d.IsDir() {
			// захист, якщо dest всередині source
			if path == destAbs {
				return filepath.SkipDir
			}
			return nil
		}

		cat := fileCategory(path)
		destDir := filepath.Join(destAbs, cat)
		return moveFile(path, destDir, logger)
	})

	if err != nil {
		logger.Printf("ERROR WalkDir: %v\n", err)
	}

	fmt.Printf("Сортування завершено! Див. '%s'\n", logFilePath)
	logger.Println("=== Sorting finished ===")
}
