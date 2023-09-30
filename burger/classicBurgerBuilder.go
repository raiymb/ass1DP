package burger

type ClassicBurgerBuilder struct{}

func (cb *ClassicBurgerBuilder) BuildBurger() *Burger {
	classicBurger := &Burger{}
	classicBurger.AddIngredient("Beef Patty")
	classicBurger.AddIngredient("Lettuce")
	classicBurger.AddIngredient("Tomato")
	classicBurger.AddIngredient("Onion")
	classicBurger.AddIngredient("Ketchup")
	classicBurger.AddIngredient("Mustard")
	return classicBurger
}
