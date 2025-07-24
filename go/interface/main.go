package main

import "fmt"

type engine interface {
	Start()
}

type gasEngine struct{}

func (g *gasEngine) Start() {
	fmt.Println("Gas Engine is starting")
}

type Car struct {
	name string
	eng  engine
}

func main() {
	var e engine
	e = &gasEngine{}

	e.Start()

	car := Car{
		name: "Gaz",
		eng:  e,
	}

	car.eng.Start()
}
