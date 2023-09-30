package burger

type VeggieBurgerBuilder struct{}

func (vb *VeggieBurgerBuilder) BuildBurger() *Burger {
	veggieBurger := &Burger{}
	veggieBurger.AddIngredient("Veggie Patty")
	veggieBurger.AddIngredient("Lettuce")
	veggieBurger.AddIngredient("Tomato")
	veggieBurger.AddIngredient("Onion")
	veggieBurger.AddIngredient("Mayonnaise")
	return veggieBurger
}
