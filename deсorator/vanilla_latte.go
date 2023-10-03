package deсorator

type VanillaLatte struct {
	CoffeeDecorator
}

func (v *VanillaLatte) GetPrice() int {
	coffeePrice := v.Coffee.GetPrice()
	return coffeePrice + 4
}
