package main

import (
	"fmt"
	"os"
	"syscall"
)

func main() {
	pid, _, errno := syscall.Syscall(syscall.SYS_FORK, 0, 0, 0)
	if errno != 0 {
		fmt.Println("Ошибка:", errno)
		os.Exit(1)
	}

	if pid == 0 {
		fmt.Println("Дочерний процесс:", os.Getpid())
	} else {
		fmt.Println("Родительский процесс:", os.Getpid(), "запустил процесс с PID", pid)
	}
}
