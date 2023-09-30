package main

import (
	"ass1/burger"
	"fmt"
)

func main() {
	classicBuilder := &burger.ClassicBurgerBuilder{}
	classicContext := burger.NewBurgerContext(classicBuilder)
	classicBurger := classicContext.CreateBurger()

	fmt.Println("Classic Burger Ingredients:")
	fmt.Println(classicBurger.GetIngredients())

	veggieBuilder := &burger.VeggieBurgerBuilder{}
	veggieContext := burger.NewBurgerContext(veggieBuilder)
	veggieBurger := veggieContext.CreateBurger()

	fmt.Println("\nVeggie Burger Ingredients:")
	fmt.Println(veggieBurger.GetIngredients())
}
