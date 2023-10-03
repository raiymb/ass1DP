package decorator

type Cappuccino struct {
	CoffeeDecorator
}

func (c *Cappuccino) GetPrice() int {
	coffeePrice := c.Coffee.GetPrice()
	return coffeePrice + 3
}
