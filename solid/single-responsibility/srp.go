package main

import (
	"fmt"
	"os"
)

type FileHandler struct{}

func (fh *FileHandler) SaveToFile(data, filename string) error {
	file, err := os.Create(filename)

	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.WriteString(data)
	return err
}

type Logger struct{}

func (l *Logger) Log(message string) {
	fmt.Println("Log:", message)
}

func main() {
	fh := FileHandler{}
	logger := Logger{}
	if err := fh.SaveToFile("Hello, SOLID!", "solid.txt"); err == nil {
		logger.Log("File saved successfully")
	}
}
