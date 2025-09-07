package main

import (
	"fmt"
	"math"
)

func main() {
	task1()
	task2()
}

// ---task 1---///
type Item struct {
	Name  *string
	Price int
}

func updateItem(item *Item, newName string, newPrice int) {
	*item = Item{
		Name:  &newName,
		Price: newPrice,
	}
	// item.Name = &newName
	// item.Price = newPrice
}

func (i *Item) String() string {
	return fmt.Sprintf("Название: %s / цена: %d", *i.Name, i.Price)
}

func task1() {
	defaulName := "Книга 1"
	item := &Item{
		Name:  &defaulName,
		Price: 15,
	}

	secondItem := &Item{}
	updateItem(secondItem, "Книга 2", 10)

	fmt.Println(secondItem)
	fmt.Println(item)
}

//---task 1---///

// ---task 2---///
type Shape interface {
	Area() float64
}

type Circle struct {
	R float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.R * c.R
}
func (c *Circle) SetR(r float64) {
	c.R = r
}

func GetArea(s Shape) float64 {
	return s.Area()
}

func task2() {
	c := &Circle{R: 10}
	fmt.Println(GetArea(c))
	c.SetR(15)
	fmt.Println(GetArea(c))
}

//---task 2---///
