package main

import (
	"fmt"
	"os"
	"syscall"
	"time"
)

/*
Create a new process indepadent of the parent process
*/

func main() {
	pid, err := syscall.ForkExec("/proc/self/exe", os.Args, &syscall.ProcAttr{
		Files: []uintptr{0, 1, 2}, // stdin, stdout, stderr
	})
	if err != nil {
		panic("Не удалось форкнуть процесс: " + err.Error())
	}

	if pid > 0 {
		fmt.Println("Процесс запущен в фоне, PID:", pid)
		os.Exit(0)
	}

	syscall.Setsid()
	os.Stdout.Close()
	os.Stderr.Close()
	os.Stdin.Close()

	for {
		time.Sleep(10 * time.Second)
		fmt.Println("Демон работает...")
	}
}
