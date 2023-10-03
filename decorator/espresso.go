package decorator

type Espresso struct{}

func (e *Espresso) GetPrice() int {
	return 5
}

type CoffeeDecorator struct {
	Coffee ICoffee
}
