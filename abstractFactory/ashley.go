package abstractFactory

import "fmt"

type AshleyFactory struct{}

func (f *AshleyFactory) CreateChair() IChair {
	return &AshleyChair{}
}

func (f *AshleyFactory) CreateTable() ITable {
	return &AshleyTable{}
}

type AshleyChair struct{}

func (c *AshleyChair) Sit() {
	fmt.Println("Sitting on an Ashley chair")
}

type AshleyTable struct{}

func (t *AshleyTable) Eat() {
	fmt.Println("Eating on an Ashley table")
}
