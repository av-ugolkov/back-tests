package main

import (
	"fmt"
	"syscall"
)

func main() {
	args := []string{"/bin/ls", "-l"}
	env := []string{"PATH=/bin"}

	pid, err := syscall.ForkExec(args[0], args, &syscall.ProcAttr{
		Dir:   "/",
		Env:   env,
		Files: []uintptr{0, 1, 2}, // stdin, stdout, stderr
	})

	if err != nil {
		panic(err)
	}

	fmt.Println("Процесс запущен, PID:", pid)
}
