package main

import (
	"fmt"
	"os"
	"strconv"
	"syscall"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Использование: go run kill.go <pid>")
		os.Exit(1)
	}

	pid, _ := strconv.Atoi(os.Args[1])
	err := syscall.Kill(pid, syscall.SIGKILL)
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println("Процесс", pid, "уничтожен")
	}
}
