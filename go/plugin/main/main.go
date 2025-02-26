package main

import (
	"fmt"
	"os"
	"plugin"
)

type Sayer interface {
	Says() string
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("usage: run main/main.go animal")
		os.Exit(1)
	}

	name := os.Args[1]
	module := fmt.Sprintf("./%s/%s.so", name, name)

	_, err := os.Stat(module)
	if os.IsNotExist(err) {
		fmt.Println("can't find an animal named", name)
		os.Exit(1)
	}

	p, err := plugin.Open(module)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	symbol, err := p.Lookup("Animal")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	animal, ok := symbol.(Sayer)
	if !ok {
		fmt.Println("that's not an Sayer")
		os.Exit(1)
	}

	fmt.Printf("A %s says: %q\n", name, animal.Says())
}
