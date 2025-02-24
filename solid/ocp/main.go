package main

import "fmt"

type DiscountStrategy interface {
	ApplyDiscount(price float64) float64
}

type RegularCustomer struct{}

func (rc RegularCustomer) ApplyDiscount(price float64) float64 {
	return price * 0.9
}

type PremiumCustomer struct{}

func (pc PremiumCustomer) ApplyDiscount(price float64) float64 {
	return price * 0.8
}

type NewCustomer struct{}

func (nc NewCustomer) ApplyDiscount(price float64) float64 {
	return price * 0.95
}

func main() {
	var customer DiscountStrategy = PremiumCustomer{}
	fmt.Println("Discounted price:", customer.ApplyDiscount(100))
}
