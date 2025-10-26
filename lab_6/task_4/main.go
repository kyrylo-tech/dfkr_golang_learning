package main

import "fmt"

type Logger interface {
	Log(message string)
}

type ConsoleLogger struct {
	Prefix string
}

func (c ConsoleLogger) Log(message string) {
	fmt.Printf("[%s] %s\n", c.Prefix, message)
}

type FileLogger struct {
	Filename string
}

func (f FileLogger) Log(message string) {
	fmt.Printf("[%s] %s\n", f.Filename, message)
}

func processLog(l Logger, msg string) {
	l.Log(msg)
}

func main() {
	cl := ConsoleLogger{Prefix: "WARN"}
	fl := FileLogger{Filename: "app.log"}

	cl.Log("Система стабільна.")
	fl.Log("Логування в файл розпочато.")

	processLog(cl, "Попередження про навантаження.")
	processLog(fl, "Дані записані у файл.")
}
