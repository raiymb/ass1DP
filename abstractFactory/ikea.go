package abstractFactory

import "fmt"

type IkeaFactory struct{}

func (f *IkeaFactory) CreateChair() IChair {
	return &IkeaChair{}
}

func (f *IkeaFactory) CreateTable() ITable {
	return &IkeaTable{}
}

type IkeaChair struct{}

func (c *IkeaChair) Sit() {
	fmt.Println("Sitting on an Ikea chair")
}

type IkeaTable struct{}

func (t *IkeaTable) Eat() {
	fmt.Println("Eating on an Ikea table")
}
