package observer

import "fmt"

type Customer struct {
	Name string
}

func NewCustomer(name string) *Customer {
	return &Customer{Name: name}
}

func (c *Customer) Update(burger *Burger) {
	fmt.Printf("Customer %s received a %s!\n", c.Name, burger)
}
