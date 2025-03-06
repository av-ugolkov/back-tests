package main

import (
	"os"
	"syscall"
)

func main() {
	err := syscall.Exec("/bin/ls", []string{"ls", "-l"}, os.Environ())
	if err != nil {
		panic(err)
	}
}
